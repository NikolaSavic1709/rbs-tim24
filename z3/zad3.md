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
       
   -
