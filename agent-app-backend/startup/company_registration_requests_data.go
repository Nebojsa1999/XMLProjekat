package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var companyRegistrationRequests = []*domain.CompanyRegistrationRequest{
	{
		Id:                 getObjectId("123a0cc3a34d25d8567f9f01"),
		OwnerId:            getObjectId("123a0cc3a34d25d8567f9f04"),
		Status:             enums.Accepted,
		ReasonForRejection: "",
		Name:               "Continental",
		Address:            "Narodnih heroja 3, Novi Sad",
		Email:              "ljudskiresursi.ns@continental.com",
		Phone:              "021/436-907",
		AreaOfWork: "Kompanija razvija širok spektar proizvoda kao što su instrument table, radari, " +
			"kamere, kontrolne jedinice za daljinsko praćenje pozicije i parametara vozila, pristupni " +
			"sistemi zasnovani na pristupu vozilu pomoću pametnih telefona ili daljinskog otključavanja, " +
			"upravljački sklop za kontrolu električnih sistema u vozilu, infrastruktura web baziranog " +
			"servisa, mobilne aplikacije i glavni inovativni projekat za visoko autonomnu vožnju za " +
			"putničke automobile i komercijalna vozila.",
		Description: "Continental je 2017. godine u Novom Sadu otvorio Istraživačko razvojni centar u " +
			"kome preko 600 inženjera radi na razvoju najinovativnijih tehnologija za automobilsku " +
			"industriju. Sa atraktivnim projektima iz oblasti elektronskih uređaja za kabinu vozila i " +
			"sistemima i senzorima za autonomnu vožnju, kompanija je za manje od četiri godine rada izrasla" +
			" u jednog od najvećih poslodavaca u inženjerskom sektoru automobilskoj industriji u zemlji.",
		WorkCulture: "Parking i biciklana: Svakom zaposlenom je obezbeđena kartica za parking. Ukoliko " +
			"dolaziš na posao biciklom imamo i biciklanu u kojoj možeš odložiti svoj bicikl.\n\n" +
			"Fleksibilno radno vreme: Jedna od pogodnosti je sloboda da odredite kada ćeš raditi. Naše " +
			"osnovno vreme je od 9:30 do 15:30, a ostatak radnog vremena možeš organizovati po sopstvenom " +
			"nahođenju.\n\n" +
			"Rad od kuce i placena naknada za rad od kuce: U roku od mesec dana dobijaš četvorodnevni " +
			"mobilni rad. Te dane možeš koristiti po sopstvenom nahođenju. Usled novonastale situacije " +
			"zaposleni mogu odabrati da rade iz kancelarije ili od kuce bez ograničenja dana. Za vreme " +
			"rada van kancelarija isplaćuje za dnevna naknada za rad.\n\n" +
			"Naknada za preporuke: Uvek smo u potrazi za kvalitetnim ljudima koji bi doprineli razvoju " +
			"našeg tima i kompanije u celini. Ako poznaješ nekoga ko je po vašem mišljenju pogodan kandidat" +
			", možeš ga uputiti u HR i steći pravo na naknadu za svoj trud i predstavljanje kompanije u " +
			"dobrom svetlu.\n\n" +
			"Godišnji odmor: Imaš pravo na 24 radna dana godišnjeg odmora po godini. Od 2022. godine " +
			"povećava se broj dana godišnjeg odmora po osnovu ukupnog radnog staža u kompaniji i to od 3 " +
			"do 5 godina staža u našoj kompaniji - 1 dodatni dan, od 5 do 10 godina - 1 dodatni dan i " +
			"preko 10 godina – 1 dodatni dan.\n\n" +
			"Privatno zdravstveno osiguranje: Naši zaposleni takođe imaju obezbeđeno i dodatno privatno " +
			"zdravstveno osiguranje. Na ovaj način je proširen spektar zdravstvene zaštite svakog " +
			"zaposlenog.\n\n" +
			"Mentorstvo: Ukoliko si naš novi član tima, naše iskusne kolege i koleginica koji će osigurati " +
			"da se upoznaš sa svojim novim radnim mestom. Imamo prostora za sva tvoja pitanja.\n\n" +
			"Podesivi stolovi i stolice: Naše kancelarije su opremljene potpuno podesivim stolovima i " +
			"stolicama. Na taj način svi naši zaposleni, bez obzira na visinu, mogu uživati u istoj " +
			"udobnosti dok rade sedeći ili stojeći.\n\n" +
			"Popusti za Continental zaposlene: Svi zaposleni Continental-a mogu ostvariti popuste u " +
			"raznim restoranima, prodavnicama, ordinacijama sa kojima kompanija ima zaključen ugovor.",
	},
	{
		Id:                 getObjectId("123a0cc3a34d25d8567f9f02"),
		OwnerId:            getObjectId("123a0cc3a34d25d8567f9f01"),
		Status:             enums.Pending,
		ReasonForRejection: "",
		Name:               "Itekako",
		Address: "Terazije 23, Beograd. Ulaz u zgradu je odmah pored prodavnice Deichmann, " +
			"na ulazu u Nušićevu, VII sprat.",
		Email: "office@itekako.com",
		Phone: "011/407-8146",
		AreaOfWork: "Atlas Copco je klijent iz Švedske za koga razvijamo upravljački interfejs za " +
			"njihove uređaje radeći direktan razvoj na hardveru, kao i alati za koordinaciju rada " +
			"kompletne mreže prodajnih mesta. Tehnologije koje se koriste su C++, Python, React " +
			"i SharePoint.\n\n" +
			"Unity je svetski poznata kompanija za koju razvijamo različite veb alate. Unity se bavi " +
			"cross-platform game engine-om koji se koristi za 2D i 3D video igre i različite simulacije " +
			"za desktop i mobile. Tehnologije koje se koriste su JavaScript (React), NodeJS (Express), " +
			"C#, Angular i PostgreSQL.\n\n" +
			"BetterCollective je danski klijent sa kojim sarađujemo već dugo, a koji posluje širom " +
			"sveta. Bave se sportskim bettingom. Na backend-u je tehnologija PHP, na frontend-u Angular" +
			", a za agregaciju koristi se graphQL.\n\n" +
			"Burda je klijent za koga razvijamo web platformu za kupovinu šivaćeg materijala i za " +
			"okupljanje zajednice čiji je zajednički hobi šivenje, a koja broji preko pola miliona " +
			"korisnika. Tehologije zastupljene na ovom projektu su Drupal 8, Magento 2, Elasticsearch, " +
			"Redis, Akeneo, PHP.\n\n" +
			"LTSE je američki startap iz San Franciska za koga razvijamo novu platformu u oblasti " +
			"finansija i  berzanskog poslovanja. Tehnologije koje se koriste su Java, PHP, React " +
			"i Nightwatch.js.",
		Description: "Itekako je beogradska firma nastala 2006. godine, kada se grupa entuzijasta okupila " +
			"sa idejom o razmeni znanja i pisanja najkvalitetnijeg koda. Glavna ideja se održala do danas" +
			", a vremenom naši glavni aduti su postali prijateljska atmosfera, kultura kontinuiranog, " +
			"svakodnevnog učenja i deljenja znanja, kao i izuzetna fleksibilnost svih aspekata poslovanja " +
			"i rada. Naš tim čini preko 120 profesionalaca, koji su diplomirani, master i doktori " +
			"prirodnih nauka- matematičari, fizičari, programeri, uz ponekog filozofa, planinara, " +
			"ljubitelja kapoeire, i mnogo muzičara, skijaša i krosfitera.",
		WorkCulture: "Fleksibilno radno vreme - ali, zaista fleksibilno!\n\n" +
			"Paket privatnog zdravstvenog osiguranja bez participacije.",
	},
}
