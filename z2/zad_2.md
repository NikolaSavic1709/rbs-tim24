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


## Zadatak B
### Osetljiva imovina

1. **Korisnički podaci - adrese, brojevi telefona, e-mail adrese...**:
   - **Izloženost**: Dostupni zaposlenima koji rade na korisničkoj podršci i administratorima
   - **Bezbednosti ciljevi (CIA)**:
     - Poverljivost - zaštita ličnih i finansijskih podataka korisnika
      - Integritet - osiguravanje tačnosti detalja rezervacija
      -  Dostupnost - osiguravanje pristupa podacima kada je potrebno
   - **Uticaj oštećenja**: Gubitak poverenja korisnika, kazne zbog kršenja privatnosti podataka, gubitak reputacije kompanije, smanjenje broja korisnika

2. **Finansijski podaci - brojevi kartica, informacije o računima i transakcijama...**:
   - **Izloženost**: Dostupni zaposlenima u finansijskom sektoru eksternog platnog provajdera ili zaposlenima u finansijskom sektoru same kompanije ukoliko ona poseduje interni sistem za plaćanje
   - **Bezbednosti ciljevi (CIA)**:
      - Poverljivost - zaštita finansijskih podataka i transakcija
      - Integritet - osiguravanje tačnosti finansijskih zapisa
      -  Dostupnost - osiguravanje tačnosti finansijskih zapisa
   - **Uticaj oštećenja**: Finansijski gubitak kompanije i korisnika, krađa identiteta, kazne zbog kršenja privatnosti podataka, nepoverenje korisnika
3. **Intelektualna svojina i poslovni procesi - autorska prava, načini poslovanja, poslovne strategije, analize tržišta, tajne informacije...**:
   - **Izloženost**: Dostupni određenim timovima unutar kompanije zaduženim za istraživanje, razvoj, marketing i strategiju poslovanja
   - **Bezbednosti ciljevi (CIA)**:
     - Poverljivost - zaštita vlasničkih informacija i poslovnih tajni
      - Integritet - osiguravanje tačnosti i pouzdanosti IP
      -  Dostupnost - osiguravanje pristupa IP za poslovne operacije
   - **Uticaj oštećenja**: Gubitak prednosti nad konkurencijom, smanjenje prihoda, gubitak pravnih sporova
4. **Sistemska infrastruktura - serveri, baze podataka, mrežne komponente...**:
   - **Izloženost**: Dostupna sistemskim administratorima i IT osoblju
   - **Bezbednosti ciljevi (CIA)**:
     - Poverljivost - zaštita osetljivih informacija koje se čuvaju u IT sistemima
      - Integritet - osiguravanje integriteta podataka i pouzdanosti sistema
      -  Dostupnost - osiguravanje operativnosti i pristupa IT sistemima
   - **Uticaj oštećenja**: Prekid usluga, gubitak podataka i reputacije, smanjena produktivnost kompanije

S obzirom da korporacija ima tri glavna sedišta u Hong Kongu, Bostonu i Londonu, postoje različiti načini za upravljanje osetljivom imovinom kao i zakoni koji doprinose smanjenju rizika od sajber napada.

1. **Hong Kong**
    -  Tokom protekle decenije je imovina organizacija drastično više izložena sajber napadima. Konstantno se predlažu nove izmene zakona kako bi se ovi problemi rešili. Za kompanije iz sektora bankarstva, hartija od vrednosti i osiguranja postoje određena regulatorna tela poput HKEX, SFC, HKMA i IA koja očekuju efikasno upravljanje poverljivim informacijama i adekvatne reakcije u slučaju incidenata koji ugrožavaju poverljivost
2. **Boston**
   - Savezna država Masačusets propisuje standarde zaštite ličnih informacija stanovnika zakonima <em>Standards for the Protection of Personal information of Residents of the Commonwealth</em>, <em>Agency Privacy Rules</em>, i <em>Data Breach Notification Law</em>. Bezbedonosni ciljevi su podeljeni na tehničke, administrativne i operativne ciljeve
3. **London**
   - Kao jedan od globalnih finansijskih centara, usled digitalizacije finansijskih usluga, London trpi veliki rizik od sajber napada. Regulatorna tela poput FCA i PRA postavljaju stroge zahteve za bezbednost podataka i upravljanje rizicima u sektoru finansija. Osim toga, GDPR donosi obaveze zaštite privatnosti podataka korisnika i građana EU, koje se takođe primenjuju na organizacije u Londonu. <em>Greater London Authority</em> (GLA) je organ odgovoran za upravljanje Londonom po pitanju transporta, urbanizma, ekonomskog razvoja i kulture. U prethodnom periodu, GLA se susreće sa raznim pretnjama na nivou sajber bezbednosti i ističu sve veću ranjivost podataka, računarskih sistema i usluga. 
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
