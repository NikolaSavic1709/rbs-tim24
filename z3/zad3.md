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
       
### Broken Access Control napadi

   - Ovi napadi predstavljaju grupu napada gde korisnici dobijaju pristup resursima na koje ne bi trebali imati pravo ili da imaju veći nivo pristupa nego što bi trebalo. Sistem u ovom slučaju ne reguliše pravilno pristup resursima na osnovu identiteta korisnika ili uloga.
     
   - Ranjivost na ove napade se može iskoristiti za sledeće svrhe:
        - neovlašćen pristup informacijama
        - integritet podataka korisnika ili sistema
        - vršenje nedozvoljenih akcija za određenu ulogu korisnika

   - Ranjivosti u softveru podložnom ovim napadima su:
        - neadekvatna autentifikacija i autorizacija
        - loša provera prava pristupa
        - zaobilaženje kontrole pristupa kroz otvorene API-je i URL-ove

   - Kontramere koje bi se trebale preduzeti kako bi se sprečili ovi napadi su:
        - implementacija jačih auth metoda (Multi-factor auth), korišćenje sigurnih algoritama za skladištenje lozinki, upotreba samo jakih lozinki
        - definisanje jasnih politika i privilegija za kontrolu pristupa
        - upotreba tokena za pristup API-jima za zaštićene resurse, stalno ažuriranje i testiranje API-ja, korišćenje enkriptovanih protokola
          
### Broken Authentication

   - Ovi napadi predstavljaju grupu napada gde podsistem za autentifikaciju nije pravilno implementiran. U ovom slučaju, podsistem za autentifikaciju ne upravlja dobro sesijama i tokenima, koristi nedovoljno sigurno lozinke i nema dodatnih algoritama zaštite. 

   - Ranjivost na ove napade se može iskoristiti za sledeće svrhe:
         - napadač može dobiti pristup osteljivim podacima, kao što su korisnički podaci, podaci vezani za finansijske kartice i račune
         - krađa identiteta
         - zloupotreba računa
         - kompromitovanje celog sistema

   - Ranjivosti u softveru podložnom ovim napadima su:
         - slabe kontrole pristupa
         - loše upravljanje sesijama i tokenima
         - loša implementacija Multi-factor Auth
         - neadekvatni algoritmi za enkripciju lozinki i tokena
     
   - Kontramere koje bi se trebale preduzeti kako bi se sprečili ovi napadi su:
         - definisanje jasnih politika i privilegija za kontrolu pristupa
         - postavljanje kratkih vremenskih perioda sesija i generisanje slučajnih sesijskih ID-eva
         - korišćenje sigurnosnih alata i algoritama za Auth

## Tim je rešio izazove:
   - XSS kategorija:
      - DOM XSS
      - Bonus Payload
   - Injection kategorija:
      - Login Admin
      - Login Jim
      - Login Bender
      - Christmas Special
      - Database Schema
      - Ephemeral Accountant
   - Sensitive Data Exposure kategorija:
      - Forgotten Developer Backup
      - Forgotten Sales Backup
      - Misplaced Signature File
      - Visual Geo Stalking
   - Improper Input Validation kategorija:
      - Poison Null Byte
      - Deluxe Fraud
   - Broken Access Control kategorija:
      - Forged Review
      - Easter Egg
   - Broken Authentication kategorija:
      - Bjoern's Favorite Pet
      - Reset Bjoern's Password
   - Broken Anti Automation kategorija:
      - CAPTCHA Bypass
   - Cryptographic Issues kategorija:
      - Nested Easter Egg
         
