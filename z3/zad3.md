## Analiza izvršenih napada i potencijalne odbrane
### Injection napadi

   - Omogućavaju napadaču ubacivanje:
     - malicionznih upita koji menjaju upit ka bazi
     - malicionznog koda u program
     - malicioznog softvera u računar

     Napad se vrši kroz ulazne tačke u softver kao što su polja za unos, URL parametri, HTTP zaglavlja i polja za otpremanje datoteka.

     Najpopularnije vrste su SQL Injection, Cross Site Scripting (XSS), OS Command Injection, Server-Side Template Injection (SSTI), HTTP Header Injection, Code Injection Attacks...
     
   - Ranjivost na ove napade se može iskoristiti za sledeće svrhe:
     - neovlašćeni pristup osetljivim podacima
     - neovlašćena izmena ili brisanje podataka
     - izmena i povećavanje sopstvenih privilegija
     - preuzimanje kontrole ili ugrožavanje celog sistema ili mreže
   
   - Ranjivosti u softveru podložnom ovim napadima su:
     -  nedostatak validacije korisničkog unosa
     -  nedostatak parametrizovanih upita u interakciji sa bazom podataka
     -  nedostatak enkodiranja i sanitizacije unosa
       
   - Primenom ovih sigurnosnih principa Injection napadi bi se mogli sprečiti


### Sensitive Data Exposure napadi

   - Ovi napadi predstavljaju grupu napada gde napadač može doći do poverljivih informacija o korisnicima, o kreditnim karticama i poslovnim tajnama. Do ovih podataka se može doći pristupanjem nezaštićenoj bazi podataka, prisluškivanjem neenkriptovane komunikacije, usled nedostatka autentifikacije i autorizacije ili nesavesnim postupanjem korisnika.

   - Ranjivost na ove napade se može iskoristiti za sledeće svrhe:
      - krađa identiteta
      - finansijska šteta
      - pristup privatnim podacima
      - opadanje ugleda kompanije i krađa poslovnih tajni
        
   - Ranjivosti u softveru podložnom ovim napadima su:
     - nedostatak enkripcije podataka
     - slab sistem za autentifikaciju
     - neadekvatna kontrola pristupa
     - nedostatak logovanja i praćenja sistema
     - socijalni inženjering

   - Kontramere koje bi se trebale preduzeti kako bi se sprečili ovi napadi su:
     - primena jakih i preporučenih algoritama za enkripciju podataka u skladištu i u tranzitu
     - zahtevanje složenijih lozinki i implementacija snažnijih metoda autentifikacije kao što su biometrijska i multifaktorska
     - definisanje jasnih politika i privilegija za kontrolu pristupa
     - postavljanje sistema za logovanje i redovno praćenje aktivnosti
     - edukacija osoblja o socijalnom inženjeringu radi prepoznavanja zlonamernih aktivnosti
       
### Broken Access Control napadi - easter egg, forged review
### Cryptographic Issues - nested easter egg
### Broken Authentication - 
