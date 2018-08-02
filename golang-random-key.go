package golang-random-key

import(
       "math/rand"
       "time"
)

func CreateRandomKey(size int) string{
     s:=""
     for i:=1;i<=size;i++{
         s+=GetSingleRandomCharacter()
     }
     return s
}

func GetSingleRandomCharacter() string{
     rand.Seed(time.Now().UnixNano())
     alphabet:="abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
     k:=string(alphabet[rand.Intn(len(alphabet))])
     return k
}


