## OSINT
### Zadatak 4 - Nigerian Prince
 - Čitanjem mejla došli smo do nekih specijalnih podataka kao što su "Senate bill 1622 
, Title 8 ; Section 308". Pretragom ovog teksta saznali smo da je za ovaj spam mejl korišćen spam mimic alat. Decode-ovanjem mejla dobili smo flag UNS{EM4IL_5P4M_AG4N?} . Na ovaj tekst možemo gledati kao leet speak, odnosno u prevedenom značenju je EMAIL_SPAM_AGAN. 
    
### Zadatak 5 - Educational Purposes Only
 - Datum osnivanja FTN-a smo našli na vikipedia (18/05/1960). 
 - Ime dekana od 1975. do 1977. smo pronašli na [istorijat-finkcije-dekana](http://www.ftn.uns.ac.rs/n539335203/istorijat-funkcije-dekan) na ftn.uns.ac.rs (Dragutin).
 - Datum otvaranja sajta FTN-a je ujedno i datum prve objave na sajtu (18/05/2005).
 - Godina osnivanja smera "Poštanski saobraćaj i telekomunikacije" je pronađena na sajtu departmana za [saobraćaj](https://saobracaj.ftn.uns.ac.rs/) pod sekcijom ISTORIJA (1999).
 - Otključavanjem arhive sa 
> unrar x -p18/05/1960Dragutin18/05/20051999 old.rar

dobili smo sliku na kojoj je flag UNS{V3RY_OLD_4RCH1V3} . 

   ### Zadatak 6 - Pixel Perfect
 - Sliku smo uploadovali na [aperisolve](https://www.aperisolve.com/) i alat je pronašao flag skriven u pikselima. Flag je UNS{PMF_5TUD3NT5_LOV3_M4TH}.
 - 
  ### Zadatak 7 - The Queen of the Ocean
- Našli smo da je pravo ime ajkule "The Queen of the Ocean" "Nukumi". Na sajtu [ocearch](https://www.ocearch.org/tracker) smo pronašli kretanje te ajkule i da je poslednji put poslala signal Apr 11, 2021, 3:03:07 PM. Flag je RC15{Apr 11, 2021, 3:03:07 PM}

 ### Zadatak 8 - Squad Game Invitation
- Analizom naloga sa imenom squidgameph1337@gmail.com pomocu sherlock project-a našli smo više naloga od kojih je jedan github. U repository-u squidgame u fajlu [index.html](https://github.com/squidgameph1337/squidgame/blob/main/index.html) nadjen je flag u u div elementu za "phone-number". Flag je NAVY{h4v3_y0u_3v3r_w4tched_!t?}
   
 ### Zadatak 9 - Maps OSINT 1
- Pronalskom kuće Save Šumanovića pronašli smo identičnu raskrsnicu. Kordinate su 45.12634206456855,  19.22924411281482. Flag je UNS{45.1263420,19.2292441}

 ### Zadatak 10 - Maps OSINT 2
- Sliku smo ubacili u image search Google-a koji nas je naveo da tražimo lokaciju na Zaovinskom jezeru. Lokacija slike je 43.87028223634826, 19.385401549856507. Flag je UNS{43.8702822,19.3854015}
