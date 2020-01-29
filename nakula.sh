URI=$1
function Initial() {
    if [ "$1" == "" ]; then
        echo "[+] Nakula adalah seorang tokoh protagonis dalam wiracarita Mahabharata." 
        echo "    Ia merupakan putra Madri dan Pandu. Ia adalah saudara kembar Sadewa dan dianggap putra Dewa Aswin, dewa tabib kembar."
        CpToTmp
        CreateApps
    else
        echo "Add the URL shellcode" 
    fi
}

function CreateApps(){
    echo -e '
package main\n
import ("fmt")\n
func main() {\n
    fmt.Println("$URI")
\n}' >> /tmp/nakula/src/apps.go 
}


function CpToTmp() {
    mkdir /tmp/nakula;
    cp -r src/ /tmp/nakula/src;
    rm /tmp/nakula/src/apps.go;
}

Initial