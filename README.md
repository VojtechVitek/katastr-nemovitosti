# Katastr nemovitostí - SOAP API pro jazyk Go

https://www.cuzk.cz/Katastr-nemovitosti/Poskytovani-udaju-z-KN/Dalkovy-pristup/Webove-sluzby-dalkoveho-pristupu.aspx

## Generace kódu
Kód je generován pomocí https://github.com/fiorix/wsdl2go

1. Stáhnout Aktualizovaný popis WSDP z https://www.cuzk.cz/Katastr-nemovitosti/Poskytovani-udaju-z-KN/Dalkovy-pristup/Webove-sluzby-dalkoveho-pristupu.aspx
2. Rozbalit archiv
3. Z výsledné složky spustit následující:

```
for i in $(find . -name *.wsdl); do ( cd $(dirname $i); wsdl2go > $(basename $i).go < ./$(basename $i); ); done
find . -name *.go -exec echo mv {} ~/go/src/github.com/VojtechVitek/katastr-nemovitosti/ \;
for i in $(ls | cut -d_ -f1); do mkdir -p $i; mv $i*.go $i/; done
```
