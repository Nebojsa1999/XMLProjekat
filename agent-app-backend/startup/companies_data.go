package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
)

var companies = []*domain.Company{
	{
		Id:      getObjectId("123a0cc3a34d25d8567f9f01"),
		OwnerId: getObjectId("123a0cc3a34d25d8567f9f04"),
		Name:    "Continental",
		Address: "Narodnih heroja 3, Novi Sad",
		Email:   "ljudskiresursi.ns@continental.com",
		Phone:   "021/436-907",
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
}
