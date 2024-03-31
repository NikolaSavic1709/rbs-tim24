## Zadatak A
### Potencijalni napadači

1. **Konkurencija**:
   - **Veština**: Srednja do visoka 
   - **Pristup sistemu**: Imaju direktan pristup samo javnim delovima sistema, pokušaće da dobiju direktan pristup na neki drugi način. Koriste plaćene hakere, unutrašnje aktere ili lične hakere.
   - **Cilj**: Prekid servisa, krađa informacija, menjanje cena kako bi dobili prednost na tržištu

2. **Hacktivisti**:
   - **Veština**: Varijabilna
   - **Pristup sistemu**: Imaju direktan pristup samo javnim delovima sistema, pokušaće da dobiju direktan pristup na neki drugi način. Koriste ranjivosti sistema ili socijalni inženjering
   - **Cilj**: Objavljivanje poverljivih podataka, promocija ideologije/političkih stavova, prekid servisa.

3. **Viskotehnološki kriminalci**:
   - **Veština**: Visoka
   - **Pristup sistemu**: Imaju direktan pristup samo javnim delovima sistema, pokušaće da dobiju direktan pristup na neki drugi način. Koriste sofisticirane ranjivosti sistema ili socijalni inženjering
   - **Cilj**: Krađa podataka o finansijama, ucena, razni vidovi prevara

4. **Unutrašnji akteri**:
   - **Veština**: Varijabilna
   - **Pristup sistemu**: Imaju direktan pristup ili znanja o sistemu koja nisu javna. Koriste pristup sistemu potreban za obavljanje posla i znanja o sistemu koja poseduju usled obavljanja posla u firmi.
   - **Cilj**: Lični profit, osveta, sabotaža

5. **Plaćeni hakeri**:
   - **Veština**: Varijabilna
   - **Pristup sistemu**: Imaju direktan pristup samo javnim delovima sistema, pokušaće da dobiju direktan pristup na neki drugi način
   - **Cilj**: Da izvrše zadatak za koji su plaćeni od strane konkurencije ili nekih drugih aktera




## Zadatak C
### Površine napada

1.  **Web aplikacija**\
**Ulazne tačke**: Glavni interfejs i forme za prijavu, registraciju, rezervaciju putovanja i smeštaja. Vrste napada: SQL injection, XSS (Cross-Site Scripting), CSRF (Cross-Site Request Forgery) ...
    
2.  **Baza podataka**\
**Ulazne tačke**: SQL upiti za dobavljanje, ažuriranje, brisanje i dodavanje podataka. Osetljivi podaci o korisnicima, rezervacijama, plaćanjima i drugim ličnim informacijama mogu biti meta za napade na bazu podataka.
    
3.  **Platni sistemi**\
**Ulazne tačke**: Loša implementacija integracije sa API-em platnih sistema poput PayPal-a, Stripe-a i bankovnih API-a. Forme za unos podataka o kartici, adresi, i drugim informacijama. Napadi na platne sisteme mogu uključivati krađu kreditnih kartica, prevaru i krađu identiteta.
    
4.  **Mrežna sigurnost**\
**Ulazne tačke**: Loša mrežna infrastruktura, portovi i interfejsi na ruterima, firewall-ovima, VPN-ovima. Može biti meta za različite vrste napada kao što je DDoS.
    
5.  **Mobilne aplikacije**\
**Ulazne tačke**: API-i za komunikaciju sa serverom i lokalne datoteke.
    
6.  **Zaposleni**\
**Ulazne tačke**: Korisnička imena i lozinke zaposlenih, admin paneli. Napadi unutrašnje pretnje mogu doći od zaposlenih koji imaju pristup osetljivim informacijama ili putem socijalnog inženjeringa koji cilja zaposlene.
    
7.  **Partneri**\
**Ulazne tačke**: Forme, API-i ili drugi interfejsi koji omogućavaju pristup eksternim resursima. MegaTravel može biti izložena rizicima od napada preko svojih partnera koji imaju pristup njihovim sistemima ili podacima.
    
8.  **Fizička sigurnost**\
Fizički objekti poput data centara ili kancelarija takođe mogu biti meta za različite vrste napada, uključujući krađu opreme, neovlašćen pristup i video nadzor.
    
9.  **Pravna i regulatorna pitanja**\
**Ulazne tačke**: Ugovori, pravilnici, politike privatnosti, uslovi korišćenja. MegaTravel može biti izložena pravnim i regulatornim rizicima, kao što su kršenje zakona o zaštiti podataka ili sporovi sa korisnicima i regulatorima.
