/*******************************************************************************
 * Solvers - an example of single threaded Jacobi and Gauss-Seidel solvers
 * as well as a multi-threaded Jacobi solver.
 *******************************************************************************
*/

package main

/* Imports */

import (
        "fmt"
)

/* Type definitions */

func main() {
    var (
         ans string;
    )
    fmt.Printf("Please select a test to run\n1 : run a single threaded Jacobi solver"+
               "\n2 : run a single threaded Gauss-Seidel solver\n"+
               "3 : run a multithreaded Jacobi solver\n"+
               "Q : quit\n\n");
    for true {
        fmt.Scanf("%s\n", &ans);
        if ans == "Q" {
            break;
        }
        fmt.Printf("Input: (%s)\n",ans);
    }
    fmt.Printf("Bye bye\n");
}

