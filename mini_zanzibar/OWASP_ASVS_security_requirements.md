### OWASP Application Security Verification Standard

- Na prvom mestu su ispoštovana sva tri ključna elementa za dobru zaštitu podataka: Confidentiality, Integrity and Availability (CIA)

 1. **HTTPS komunikacija:** 
 - Korišćenje sertifikata za backend je dobra praksa koja obezbeđuje bezbednu komunikaciju. 
 - Korišćen je self-signed sertifikat za Nginx, koji se generise prilikom kreiranja slike za docker kontejner.
 - Korišćenje odvojenog sertifikata za komunikaciju između baze i backend-a takođe predstavlja dobru praksu.
 - Po ASVS standardu v4.0.3-9.1.1 zabranjuje se korišćenje HTTP komunikacije ukoliko nije moguće uspostaviti HTTPS vezu što je ispoštovano.
 - Po ASVS standardu v4.0.3-9.1.2 treba koristiti aktuelne verzije algoritama, šifara i protokola. Ovo je ispoštovano. (RSA 4096 with SHA-256)
 - Po ASVS standardu v4.0.3-9.1.3 treba verifikovati da su starije verzije onemogućene. Ovo je ispoštovano.

 2. **Rate limit:** 
 - Implementacija mehanizma zaštite od "spamovanja" predstavlja dobru praksu.
 - Korišćenje limita na serverskoj strani je ključna za efikasnu zaštitu.
 - Po ASVS standardu v4.0.3-2.2.1  potrebno je implementirati restrikciju broja upita i Rate Limit se navodi kao koristan alat. Ovo je ispoštovano na nivou Nginx reverse proxy-a.
 - Limit je postavljen na 10 zahteva po sekunde a dodatno je omogućen i burst od 5 zahteva.

 3. **Validacija podataka:** 
  - Sprečavanje Injection, XSS i Path Traversal napada su ključne bezbednosne kontrole. 
  -  Validacija podataka prema najboljim praksama doprinosi održavanju bezbednog koda. 
  -  Po ASVS standardu v4.0.3-5.#.# potrebno je sanitizovati sve inpute naše aplikacije na frontu. Na backendu su dodatne provere realizovane korišćenjem REGEX-a za forme ACL-ova i Namespace-ova.
  
4.   **Monitoring uz logove:**
  - Potrebno je razviti strategiju za efikasno praćenje logova sistema i odgovoriti na incidente u realnom vremenu.
  - Po ASVS standardu v4.0.3-7.1.1,7.1.2 aplikacije ne treba da loguje kredencijale ili druge osetljive podatke. Ovo je ispoštovano.
  - Po ASVS standardu v4.0.3-7.1.3 aplikacija treba logovati podatke relevantne za sigurnost sistema kao što su uspešna i neuspešna prijava, nevalidni inputi, pristup resursima bez adekvatnih privilegije. Ovo je ispoštovano.
  - Po ASVS standardu v4.0.3-7.1.4 log treba da sadrži podatke na osnovu kojih je moguće rekreirati i detaljno analizirati tok događaja. Ovo je ispoštovano upotrebom timestamp-a i odgovarajućeg log level-a.
  - Logovanje se vrši na dva nivoa: sa Nginx-a i go aplikacije. Logovi koji opisuju zahteve ka Nginx-u su dodatno enkriptovani zbog zloupotrebe podataka.
  
5. **Autentifikacija i autorizacija:**
  - Zanzibar aplikaciju može koristiti samo autentifikovan korisnik. Prilikom provere ACL-a korisnik može da vidi samo svoje ACL-ove ili ACL-ove Namespace gde ima relaciju "owner". Samo korisnik sa relacijom "owner" može da doda novi ACL za taj Namespace. Realizovana je i zaštita pristupa drugim komponentama kojima nemaju potrebe pristupati. 
  - Po ASVS standardu v4.0.3-1.4.1 svaki element sistema treba da zahteva odgovarajuća prava pristupa da bi se mogli koristiti. Ovo je ispoštovano. Za konekciju ka svakoj bazi podataka je potrebna lozinka, dok su podaci zaštićeni zavisno od korisnika koji je njihov vlasnik (ili ima durog pravo pristupa). 
  - Po ASVS standardu v4.0.3-1.4.5 sistem treba da proverava da li korisnik ima privilegije da pristupi nekom resursu, a ne samo rolu koju korisnik ima. Naša aplikacija ovo podržava.

6. **Zaštita osteljivih podataka pri keširanju**
  - Za go aplikaciju korišćen je nginx-a kao i load balancer. Ali se vodi računa o zaštiti podataka prilikom keširanja.  
  - Po ASVS standardu v4.0.3-8.1.1 potrebno je zaštiti osetljive podatke od keširanja u komponentama kao što su Load balancer. Ovo je ispoštovano.

7. **Build i Deploy**
  - Naš sistem je ispoštovao sve najbolje prakse za dobar i siguran build i deploy. Koriste se environment varijable za osetljive podatke i moguće je lako isključiti i ponovo pokrenuti sistem usled kvarova ili promena.
  - Po ASVS standardu v4.0.3-14.1.1, v4.0.3-14.1.2, v4.0.3-14.1.3 potrebno je zaštiti osetljive podatke i omogućiti kvalitetan buold i deploy uz najbolje prakse. Ovo je ispoštovano.

Ovo su neke dodatne bezbednosne kontrole, a naravno koristili smo i osnove zaštite podataka.
