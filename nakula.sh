echo "[+] Nakula adalah seorang tokoh protagonis dalam wiracarita Mahabharata." 
echo "    Ia merupakan putra Madri dan Pandu. Ia adalah saudara kembar Sadewa dan dianggap putra Dewa Aswin, dewa tabib kembar."

function CreateApps(){
    echo -e '
package main\n
import ("fmt")\n
func main() {\n
    fmt.Println("$1")
\n}' >> /tmp/nakula/src/apps.go 
}


function CpToTmp() {
    mkdir /tmp/nakula;
    cp -r src/ /tmp/nakula/src;
    rm /tmp/nakula/src/apps.go;
}


CpToTmp
CreateApps
