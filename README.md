## Generace kódu
1. Stáhnout Aktualizovaný popis WSDP z https://www.cuzk.cz/Katastr-nemovitosti/Poskytovani-udaju-z-KN/Dalkovy-pristup/Webove-sluzby-dalkoveho-pristupu.aspx
2. Rozbalit archiv
3. A z nové složky spustit pomocí https://github.com/fiorix/wsdl2go následující:

```
for i in $(find . -name *.wsdl); do ( cd $(dirname $i); wsdl2go > $(basename $i).go < ./$(basename $i); ); done
find . -name *.go -exec echo mv {} ~/go/src/github.com/VojtechVitek/katastr-nemovitosti/ \;
for i in $(ls | cut -d_ -f1); do mkdir -p $i; mv $i*.go $i/; done
```
