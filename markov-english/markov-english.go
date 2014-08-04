/******************************************************************************
 * markov-english
 *
 * Takes a file of words, creates a model of the underlying language (okay,
 * a model of how the words look like really), and creates words that look like
 * they are native words. Three mandatory arguments: the dictionary, the word
 * length and the score (see README). An optional argument provides a set of
 * letters to be used in the word. In that case, length becomes the minimum
 * length.
 *****************************************************************************/

package main

/*
 * Arguments/Parameters
 *
 * -d <dictionary file> - mandatory
 * -l <length>          - mandatory
 * -s <score>           - mandatory
 * -f <letters>         - optional
 *
 * The score is the -10*log10(Prob(first letter))+sum of -10*log10(P(trans))+
 * -10*log10(P(final)). The higher the score, the less likely the word.
 */

import (
        "io/ioutil"
        "fmt"
        "os"
        "strings"
        "math"
)

type mfloat map[byte]float64
type mmfloat map[byte]mfloat

func readfile(fname string) []string {
/* Read the file and split the strings into an array. Also, suppresses
   empty strings and converts everything to lowercase.
*/
    barray, ferror := ioutil.ReadFile(fname);
    if (ferror != nil) {
        fmt.Println("readfile: cannot open ", fname);
        fmt.Println(ferror)
        os.Exit(-1);
    }
    /* converts the array, discards the empty strings and lowercase
       everything
    */
    sarray := strings.Split(string(barray),"\n");
    scarray := make([]string,len(sarray));
    j := 0;
    for i:=0 ; i<len(sarray) ;i++  {
        if(len(sarray[i])>0) {
            scarray[j]=strings.ToLower(sarray[i]);
            j++
        }
    }
    return scarray[0:j]
}

func entryproba(dict *[]string) mfloat {
/* takes a list of strings and returns the probability distribution
   of the first letter
*/
   var v int;
   var lcount map[byte]int;
   var eproba mfloat;
   var deno float64;
   var b byte;
   var w string;
   lcount=make(map[byte]int);
   for _,w = range *dict {
       _,ok := lcount[w[0]];
       if !ok {
           lcount[w[0]]=0;
       }
       lcount[w[0]]++;
   }
   deno=float64(len(*dict));
   eproba=make(mfloat);
   for b,v = range lcount {
      eproba[b]=float64(v)/deno;
   }
   return eproba;
}

func transiproba(dict *[]string) mmfloat {
/* Returns the transition probabilities */
    var i,s int;
    var w string;
    var b1,b2 byte;
    var tproba mmfloat;
    var lcount map[byte]map[byte]int;
    lcount=make(map[byte]map[byte]int);
    for _,w = range *dict {
        for i=0 ; i<len(w)-1 ; i++ {
            b1=w[i];
            b2=w[i+1];
            _,ok := lcount[b1];
            if !ok {
                lcount[b1]=make(map[byte]int);
                lcount[b1][b2]=0;
            }
            lcount[b1][b2]++;
        }
    }
    /* We have counted all the transitions, let's transform
       the counts into probabilities. */
    tproba=make(mmfloat);
    for b1,_ = range lcount {
        s=0;
        for _,i = range lcount[b1] {
            s=s+i;
        }
        tproba[b1]=make(mfloat);
        for b2,i = range lcount[b1] {
            tproba[b1][b2]=float64(i)/float64(s);
        }
    }
    return tproba;;
}

func finishproba(dict *[]string) mfloat {
/* Return the finish probability distribution, that is the
   probability that a given letter finishes a word.
*/
    var w string;
    var b byte;
    var tcount map[byte]int;
    var s,i int;
    var fproba mfloat;
    tcount=make(map[byte]int);
    fproba=make(mfloat);
    s=0;
    for _,w = range *dict {
       b=w[len(w)-1];
       _,ok := tcount[b];
       if !ok {
           tcount[b]=0;
       }
       tcount[b]++;
       s++;
    }
    for b,i = range tcount {
        fproba[b]=float64(i)/float64(s);
    }
  return fproba;
}

func buildbodyword(candword string, fproba mfloat, tproba mmfloat, score float64, n int) (string,bool) {
/* This function calls the word building sequence. */
    var ll byte;
    if score < 0.0 {
    /* The word is not probable enough, let's abort */
        return "", false;
    }
    if n == 0 {
    /* No more letters, let's check if the end transition is plausible */
        ll=candword[len(candword)-1]
        if score+10*math.Log10(fproba[ll]) > 0.0 {
        /* Yes, the candidate is plausible */
            return candword, true;
        } else {
        /* Unfortunately, no */
            return "", false;
        }
    }
    /* We are in the middle of the word */
    ll=candword[len(candword)-1];
    for l,p := range tproba[ll] {
    /* For each possible letter in the transition, let's try the transition until
       we find something that works.
    */
        nextcandword:=candword+string(l);
        trycandword,worked := buildbodyword(nextcandword, fproba, tproba, score+10*math.Log10(p),
                                            n-1);
        if worked {
        /* This word works */
            return trycandword, true;
        }
    }
    /* If we are here, no transition worked correctly. */
    return "", false;
}

func buildword(eproba,fproba mfloat, tproba mmfloat, score float64, n int) string {
/* This function builds the word. The arguments are
   eproba, fproba, tproba - the probabilities as calculated.
   score - the maximum score for the word. The likelier the proba, the lower
           the score.
   n - the lenght of the word.
*/
   for b,p := range eproba {
       trycandword,worked := buildbodyword(string(b), fproba, tproba, score+10*math.Log10(p),
                                           n-1);
       if worked {
       /* This word worked. Let's return it. */
           return trycandword;
       }
   }
   /* Nothing worked, return an empty string. */
   return "";
}

func main() {
    var s1 string;
    var n int;
    sarray := readfile("dict.txt");
    eproba := entryproba(&sarray);
    tproba := transiproba(&sarray);
    fproba := finishproba(&sarray);
    for n = 4 ; n < 30; n++ {
       s1=buildword(eproba,fproba,tproba, float64(10*n), n);
       if s1 == "" {
           s1=":-(";
       }
       fmt.Printf("Length: %2d, Word : %s\n",n,s1);
    }
}

