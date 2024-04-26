## Analiza koda projekta po izboru
### Opis projekta
 - Projekat je implementacija SCADA sistema kao veb aplikacije korišćenjem ASP.NET frameworka. Aplikacija omogućava korisnicima nadzor i upravljanje industrijskim sistemima. Podržava dodavanje, uklanjanje i prikaz analognih/digitalnih tagova. Za svaki tag moguće je definisati alarm koji se aktivira na određenoj vrednosti. Sistem omogućava povezivanje na RTU (Real-Time Unit), koji šalje podatke na određeni ulazni tag. Korisnicima je omogućen lak pregled svih podataka, pregled trendinga (obeležja koja se brzo menjaju i koja imaju veliki prioritet), kao i generisanje izveštaja. Ideja je integrisati napredne funkcionalnosti kako bi se korisnicima pružili efikasni alati za optimizaciju industrijskih procesa.
### Članovi tima
  - Nikola Savić
  - Jovan Najdovski
  - Uroš Stanić
### Pronađeni defekti

  #### Statička automatska analiza koda

   - Isključena zaštita od CSRF (Cross-Site Request Forgery) napada što može dovesti do toga da treća strana iskoristi autorizaciju korisnika i izvrši izmenjen zahtev umesto njega. Rešava se uvođenjem CSRF tokena (uglavnom jednokratan) ili proveravanjem odakle je stigao zahtev.
     
     ```
     [HttpPost]
        public ActionResult<Alarm> NewAlarm(AlarmCreateDTO alarmCreateDTO)
     ```
     Potrebno je dodati atribut:
     ```
     [ValidateAntiForgeryToken]
     ```
     na metode ASP.NET MVC controller-a
     
  #### Ručna analiza koda
   - Nemamo pametnu zaštitu od SQL injection i XSS napada na nivou svih kontrolera. Korišćenjem sanitizacije i dodatnih validacija ovi problemi bi se rešili.

   - Generisanje random vrednosti
     - Umesto korišćenja System.Random klase, bolji pristup je upotreba RNGCryptoServiceProvider klase iz System.Cryptography namespace-a.
       
   - Validacija polja
     - Validacija pri dodavanju novog Alarma. Korisnički unos za Priority i Type se proverava samo za nekoliko određenih vrednosti, a sve ostalo se tretira kao greška. Detaljnija validacija bi sprečila unošenje nevalidnih ili zlonamernih vrednosti.
     - U metodi GetAllTriggers, ne postoji ograničenje za vremenski opseg koji korisnik može proslediti. Ovo može dovesti do preopterećenja sistema ako korisnik prosledi veliki vremenski opseg. Preporučuje se postavljanje gornjeg limita za vremenski opseg ili implementiranje paginacije kako bi se ograničio broj rezultata koji se vraćaju.
     - Validacija polja kao što su email adresa, lozinka, adresa, broj telefona je ranjiva na specijalne karaktere i može biti meta XSS i SQL injection napada. Ovo se može sprečiti poboljšavanjem regularnih izraza i dodatnim proverama.
     - Pri validaciji id polja, nedostaje gornja granica što može rezultovati problemima sa performansama i servera i baze podataka usled slanja ogromnih vrednosti. Ovo se može sprečiti dodavanjem gornje granice u konsultaciji sa ostalim učesnicima u razvoju softvera.
      
  - XML injection napadi
    - Iako koristimo XElement koji je proveren protiv XML injection napada trebamo izvršiti dodatne provere tagova koji se dodaju u xml konfiguraciju
      ```
      XElement newElement = new XElement(xmlTag,
          new XElement("Id", tag.Id),
          new XElement("Description", tag.Description),
          new XElement("Value", tag.Values.Count==0? "/" : tag.Values[tag.Values.Count-1].Value),
          new XElement("Timestamp", tag.Values.Count == 0 ? "/" : tag.Values[tag.Values.Count - 1].Date.ToString("HH:mm:ss:fff dd.MM.yyyy"))
      );
      ```
  
- Ostalo
    - Trenutno nema ograničenja na broj pokušaja prijave. To može omogućiti napadaču da izvodi brute force napade pokušavajući različite kombinacije korisničkih imena i lozinki. Preporuka je implementiranje mehanizma zaštite od brute force napada, kao što su blokiranje naloga nakon nekoliko neuspešnih pokušaja prijave ili postavljanje kašnjenja između pokušaja prijave.
    - Trenutno se koristi osnovna autentikacija proverom korisničkog imena i lozinke. Razmotriti dodatne mehanizme autentikacije kao što su dvo faktorska autentikacija (2FA) za poboljšanu sigurnost.
    - Metode za dobavljanje vrednosti tagova pri filteru, mogu biti meta DoS napada pošto se izvršavaju iznova i iznova na korisnički klik. Rešenje ovog problema može biti korišćenje keš memorije.
    - Potrebno dodati logovanje aktivnosti, posebno za podatke o povezivanju RTU sa sistemom. Logovanje aktivnosti može biti korisno za otkrivanje nepravilnosti ili zlonamernih aktivnosti.
    - Razmotriti upotrebu naprednijih autentifikacionih mehanizama poput Identity za ASP.NET Core za sigurniju autentifikaciju.
    - Konfiguracioni fajlovi dataConfig.xml i alarmsConfig.xml su osetljivi na neovlašćeni pristup. Potrebno koristiti adekvatne dozvole na operativnom sistemu da ograničite pristup ovim fajlovima. Moguće dodati i digitalno potpisivanje za dodatnu zaštitu.


Vreme provedeno na analizi koda je oko 4-5h. Dosta vremena provedeno je na konfiguraciji alata za statičku analizu koda.
