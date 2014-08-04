// ******************************
// * Test d'interface
// *

package main

// Les imports

import "fmt"
import "strconv"

// Les structures

type Rectangle struct {
	largeur, longueur float64
}

type Carre struct {
	largeur float64
}

type Cercle struct {
	rayon	float64
}

type Former interface {
	Forme() string
}

type Surfacer interface {
	Surface() float64
}

type Perimeter interface {
	Perimetre() float64
}

func (f Rectangle) Forme() string {
	return "rectangle"
}

func (f Carre) Forme() string {
	return "carré"
}

func (f Cercle) Forme() string {
	return "cercle"
}

func (f Rectangle) Surface() float64 {
	return f.longueur*f.largeur
}

func (f Carre) Surface() float64 {
	return f.largeur*f.largeur
}

func (f Cercle) Surface() float64 {
	return f.rayon*f.rayon*3.1415927
}

func (f Rectangle) Perimetre() float64 {
	return 2*f.longueur+2*f.largeur
}

func (f Carre) Perimetre() float64 {
	return 4*f.largeur
}

func (f Cercle) Perimetre() float64 {
	return 2*3.1415927*f.rayon
}

func main() {
	f1 := Rectangle{2,4}
	f2 := Carre{3}
	f3 := Cercle{12.0/(2*3.1415927)}

	Fo := [...]Former{f1,f2,f3}
	So := [...]Surfacer{f1,f2,f3}
	Po := [...]Perimeter{f1,f2,f3}

	l := [...]string{"La première forme ", "La deuxième forme ", "La troisième "}

	for n := 0; n < 3; n++ {
		lc := l[n] + " est un "+Fo[n].Forme()+", sa surface vaut "+strconv.FormatFloat(So[n].Surface(),'E',4,64)
		lc=lc+",et son périmètre vaut "+strconv.FormatFloat(Po[n].Perimetre(),'E',4,64)
		fmt.Printf("%s\n",lc)	
	}
	fmt.Println("La fin!")
}
