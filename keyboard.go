//testbranch#2

package keyboard

import (
      "os"
      "bufio"
      "strings"
      "strconv"
      )
      

func GetString() (string) {
    reader := bufio.NewReader(os.Stdin)
    
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    
    return input
}

func GetInt() (int) {
    reader := bufio.NewReader(os.Stdin)
    
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    
    num, _ := strconv.Atoi(input)
    
    return num
}

func GetFloat() (float64) {
    reader := bufio.NewReader(os.Stdin)
    
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    
    num, _ := strconv.ParseFloat(input, 64)
    
    return num
}
