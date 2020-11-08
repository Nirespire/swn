package name

import (
	"fmt"
	"github.com/Nirespire/swn/util"
)


type Names struct {
	Male []string
	Female []string
	Surname []string
	Place []string
}

type Culture string

const (
	Arabic   Culture = "Arabic"
	Chinese  Culture = "Chinese"
	English  Culture = "English"
	Greek    Culture = "Greek"
	Indian   Culture = "Indian"
	Japanese Culture = "Japanese"
	Latin    Culture = "Latin"
	Nigerian Culture = "Nigerian"
	Russian  Culture = "Russian"
	Spanish  Culture = "Spanish"
	Any      Culture = "Any"
)

type Gender string

const (
	Male Gender = "Male"
	Female Gender = "Female"
	None Gender = "None"
)

func getNamesByCulture(culture Culture) Names {
	if culture == Any {
		for _, name := range NamesList {
			return name
		}
	} else {
		return NamesList[culture]
	}
	return NamesList[Arabic]
}

func GetNameByCultureGender(culture Culture, gender Gender) string {
	names := getNamesByCulture(culture)

	r := util.GetNewRandom()

	place := names.Place[r.Intn(len(names.Place))]
	surname := names.Surname[r.Intn(len(names.Surname))]

	var first string
	if gender == None {
		if r.Intn(1) == 1 {
			first = names.Female[r.Intn(len(names.Female))]
		}  else {
			first = names.Male[r.Intn(len(names.Male))]
		}
	} else {
		if gender == Male {
			first = names.Male[r.Intn(len(names.Male))]
		} else {
			first = names.Female[r.Intn(len(names.Female))]
		}
	}

	return fmt.Sprintf("%s %s of %s", first, surname, place)
}


var NamesList = map[Culture]Names {
	Arabic: {
		Male: []string{"Aamir", "Ayub", "Binyamin", "Efraim", "Ibrahim", "Ilyas", "Ismail", "Jibril", "Jumanah", "Kazi", "Lut", "Matta", "Mohammed", "Mubarak", "Mustafa", "Nazir", "Rahim", "Reza", "Sharif", "Taimur", "Usman", "Yakub", "Yusuf", "Zakariya", "Zubair"},
		Female: []string{"Aisha", "Alimah", "Badia", "Bisharah", "Chanda", "Daliya", "Fatimah", "Ghania", "Halah", "Kaylah", "Khayrah", "Layla", "Mina", "Munisa", "Mysha", "Naimah", "Nissa", "Nura", "Parveen", "Rana", "Shalha", "Suhira", "Tahirah", "Yasmin", "Zulehka"},
		Surname: []string{"Abdel", "Awad", "Dahhak", "Essa", "Hanna", "Harbi", "Hassan", "Isa", "Kasim", "Katib", "Khalil", "Malik", "Mansoor", "Mazin", "Musa", "Najeeb", "Namari", "Naser", "Rahman", "Rasheed", "Saleh", "Salim", "Shadi", "Sulaiman", "Tabari"},
		Place: []string{"Adan", "Magrit", "Ahsa", "Masqat", "Andalus", "Misr", "Asmara", "Muruni", "Asqlan", "Qabis", "Baqubah", "Qina", "Basit", "Rabat", "Baysan", "Ramlah", "Baytlahm", "Riyadh", "Bursaid", "Sabtah", "Dahilah", "Salalah", "Darasalam", "Sana", "Dawhah", "Sinqit", "Ganin", "Suqutrah", "Gebal", "Sur", "Gibuti", "Tabuk", "Giddah", "Tangah", "Harmah", "Tarifah", "Hartum", "Tarrakunah", "Hibah", "Tisit", "Hims", "Uman", "Hubar", "Urdunn", "Karbala", "Wasqah", "Kut", "Yaburah", "Lacant", "Yaman"},
	},
	Chinese: {
		Male:  []string{"Aiguo", "Bohai", "Chao", "Dai", "Dawei", "Duyi", "Fa", "Fu", "Gui", "Hong", "Jianyu", "Kang", "Li", "Niu", "Peng", "Quan", "Ru", "Shen", "Shi", "Song", "Tao", "Xue", "Yi", "Yuan", "Zian"},
		Female: []string{"Biyu", "Changying", "Daiyu", "Huidai", "Huiliang", "Jia", "Jingfei", "Lan", "Liling", "Liu", "Meili", "Niu", "Peizhi", "Qiao", "Qing", "Ruolan", "Shu", "Suyin", "Ting", "Xia", "Xiaowen", "Xiulan", "Ya", "Ying", "Zhilan"},
		Surname: []string{"Bai", "Cao", "Chen", "Cui", "Ding", "Du", "Fang", "Fu", "Guo", "Han", "Hao", "Huang", "Lei", "Li", "Liang", "Liu", "Long", "Song", "Tan", "Tang", "Wang", "Wu", "Xing", "Yang", "Zhang"},
		Place: []string{"Andong", "Luzhou", "Anqing", "Ningxia", "Anshan", "Pingxiang", "Chaoyang", "Pizhou", "Chaozhou", "Qidong", "Chifeng", "Qingdao", "Dalian", "Qinghai", "Dunhuang", "Rehe", "Fengjia", "Shanxi", "Fengtian", "Taiyuan", "Fuliang", "Tengzhou", "Fushun", "Urumqi", "Gansu", "Weifang", "Ganzhou", "Wugang", "Guizhou", "Wuxi", "Hotan", "Xiamen", "Hunan", "Xian", "Jinan", "Xikang", "Jingdezhen", "Xining", "Jinxi", "Xinjiang", "Jinzhou", "Yidu", "Kunming", "Yingkou", "Liaoning", "Yuxi", "Linyi", "Zigong", "Lushun", "Zoige"},
	},
	English: {
		Male:  []string{"Adam", "Albert", "Alfred", "Allan", "Archibald", "Arthur", "Basil", "Charles", "Colin", "Donald", "Douglas", "Edgar", "Edmund", "Edward", "George", "Harold", "Henry", "Ian", "James", "John", "Lewis", "Oliver", "Philip", "Richard", "William"},
		Female: []string{"Abigail", "Anne", "Beatrice", "Blanche", "Catherine", "Charlotte", "Claire", "Eleanor", "Elizabeth", "Emily", "Emma", "Georgia", "Harriet", "Joan", "Judy", "Julia", "Lucy", "Lydia", "Margaret", "Mary", "Molly", "Nora", "Rosie", "Sarah", "Victoria"},
		Surname: []string{"Barker", "Brown", "Butler", "Carter", "Chapman", "Collins", "Cook", "Davies", "Gray", "Green", "Harris", "Jackson", "Jones", "Lloyd", "Miller", "Roberts", "Smith", "Taylor", "Thomas", "Turner", "Watson", "White", "Williams", "Wood", "Young"},
		Place: []string{"Aldington", "Kedington", "Appleton", "Latchford", "Ashdon", "Leigh", "Berwick", "Leighton", "Bramford", "Maresfield", "Brimstage", "Markshall", "Carden", "Netherpool", "Churchill", "Newton", "Clifton", "Oxton", "Colby", "Preston", "Copford", "Ridley", "Cromer", "Rochford", "Davenham", "Seaford", "Dersingham", "Selsey", "Doverdale", "Stanton", "Elsted", "Stockham", "Ferring", "Stoke", "Gissing", "Sutton", "Heydon", "Thakeham", "Holt", "Thetford", "Hunston", "Thorndon", "Hutton", "Ulting", "Inkberrow", "Upton", "Inworth", "Westhorpe", "Isfield", "Worcester"},
	},
	Greek: {
		Male:  []string{"Alexander", "Alexius", "Anastasius", "Christodoulos", "Christos", "Damian", "Dimitris", "Dysmas", "Elias", "Giorgos", "Ioannis", "Konstantinos", "Lambros", "Leonidas", "Marcos", "Miltiades", "Nestor", "Nikos", "Orestes", "Petros", "Simon", "Stavros", "Theodore", "Vassilios", "Yannis"},
		Female: []string{"Alexandra", "Amalia", "Callisto", "Charis", "Chloe", "Dorothea", "Elena", "Eudoxia", "Giada", "Helena", "Ioanna", "Lydia", "Melania", "Melissa", "Nika", "Nikolina", "Olympias", "Philippa", "Phoebe", "Sophia", "Theodora", "Valentina", "Valeria", "Yianna", "Zoe"},
		Surname: []string{"Andreas", "Argyros", "Dimitriou", "Floros", "Gavras", "Ioannidis", "Katsaros", "Kyrkos", "Leventis", "Makris", "Metaxas", "Nikolaidis", "Pallis", "Pappas", "Petrou", "Raptis", "Simonides", "Spiros", "Stavros", "Stephanidis", "Stratigos", "Terzis", "Theodorou", "Vasiliadis", "Yannakakis"},
		Place: []string{"Adramyttion", "Kallisto", "Ainos", "Katerini", "Alikarnassos", "Kithairon", "Avydos", "Kydonia", "Dakia", "Lakonia", "Dardanos", "Leros", "Dekapoli", "Lesvos", "Dodoni", "Limnos", "Efesos", "Lykia", "Efstratios", "Megara", "Elefsina", "Messene", "Ellada", "Milos", "Epidavros", "Nikaia", "Erymanthos", "Orontis", "Evripos", "Parnasos", "Gavdos", "Petro", "Gytheio", "Samos", "Ikaria", "Syros", "Ilios", "Thapsos", "Illyria", "Thessalia", "Iraia", "Thira", "Irakleio", "Thiva", "Isminos", "Varvara", "Ithaki", "Voiotia", "Kadmeia", "Vyvlos"},
	},
	Indian: {
		Male:  []string{"Amrit", "Ashok", "Chand", "Dinesh", "Gobind", "Harinder", "Jagdish", "Johar", "Kurien", "Lakshman", "Madhav", "Mahinder", "Mohal", "Narinder", "Nikhil", "Omrao", "Prasad", "Pratap", "Ranjit", "Sanjay", "Shankar", "Thakur", "Vijay", "Vipul", "Yash"},
		Female: []string{"Amala", "Asha", "Chandra", "Devika", "Esha", "Gita", "Indira", "Indrani", "Jaya", "Jayanti", "Kiri", "Lalita", "Malati", "Mira", "Mohana", "Neela", "Nita", "Rajani", "Sarala", "Sarika", "Sheela", "Sunita", "Trishna", "Usha", "Vasanta"},
		Surname: []string{"Achari", "Banerjee", "Bhatnagar", "Bose", "Chauhan", "Chopra", "Das", "Dutta", "Gupta", "Johar", "Kapoor", "Mahajan", "Malhotra", "Mehra", "Nehru", "Patil", "Rao", "Saxena", "Shah", "Sharma", "Singh", "Trivedi", "Venkatesan", "Verma", "Yadav"},
		Place: []string{"Ahmedabad", "Jaisalmer", "Alipurduar", "Jharonda", "Alubari", "Kadambur", "Anjanadri", "Kalasipalyam", "Ankleshwar", "Karnataka", "Balarika", "Kutchuhery", "Bhanuja", "Lalgola", "Bhilwada", "Mainaguri", "Brahmaghosa", "Nainital", "Bulandshahar", "Nandidurg", "Candrama", "Narayanadri", "Chalisgaon", "Panipat", "Chandragiri", "Panjagutta", "Charbagh", "Pathankot", "Chayanka", "Pathardih", "Chittorgarh", "Porbandar", "Dayabasti", "Rajasthan", "Dikpala", "Renigunta", "Ekanga", "Sewagram", "Gandhidham", "Shakurbasti", "Gollaprolu", "Siliguri", "Grahisa", "Sonepat", "Guwahati", "Teliwara", "Haridasva", "Tinpahar", "Indraprastha", "Villivakkam"},
	},
	Japanese: {
		Male:  []string{"Akira", "Daisuke", "Fukashi", "Goro", "Hiro", "Hiroya", "Hotaka", "Katsu", "Katsuto", "Keishuu", "Kyuuto", "Mikiya", "Mitsunobu", "Mitsuru", "Naruhiko", "Nobu", "Shigeo", "Shigeto", "Shou", "Shuji", "Takaharu", "Teruaki", "Tetsushi", "Tsukasa", "Yasuharu"},
		Female: []string{"Aemi", "Airi", "Ako", "Ayu", "Chikaze", "Eriko", "Hina", "Kaori", "Keiko", "Kyouka", "Mayumi", "Miho", "Namiko", "Natsu", "Nobuko", "Rei", "Ririsa", "Sakimi", "Shihoko", "Shika", "Tsukiko", "Tsuzune", "Yoriko", "Yorimi", "Yoshiko"},
		Surname: []string{"Abe", "Arakaki", "Endo", "Fujiwara", "Goto", "Ito", "Kikuchi", "Kinjo", "Kobayashi", "Koga", "Komatsu", "Maeda", "Nakamura", "Narita", "Ochi", "Oshiro", "Saito", "Sakamoto", "Sato", "Suzuki", "Takahashi", "Tanaka", "Watanabe", "Yamamoto", "Yamasaki"},
		Place: []string{"Bando", "Mitsukaido", "Chikuma", "Moriya", "Chikusei", "Nagano", "Chino", "Naka", "Hitachi", "Nakano", "Hitachinaka", "Ogi", "Hitachiomiya", "Okaya", "Hitachiota", "Omachi", "Iida", "Ryugasaki", "Iiyama", "Saku", "Ina", "Settsu", "Inashiki", "Shimotsuma", "Ishioka", "Shiojiri", "Itako", "Suwa", "Kamisu", "Suzaka", "Kasama", "Takahagi", "Kashima", "Takeo", "Kasumigaura", "Tomi", "Kitaibaraki", "Toride", "Kiyose", "Tsuchiura", "Koga", "Tsukuba", "Komagane", "Ueda", "Komoro", "Ushiku", "Matsumoto", "Yoshikawa", "Mito", "Yuki"},
	},
	Latin: {
		Male:  []string{"Agrippa", "Appius", "Aulus", "Caeso", "Decimus", "Faustus", "Gaius", "Gnaeus", "Hostus", "Lucius", "Mamercus", "Manius", "Marcus", "Mettius", "Nonus", "Numerius", "Opiter", "Paulus", "Proculus", "Publius", "Quintus", "Servius", "Tiberius", "Titus", "Volescus"},
		Female: []string{"Appia", "Aula", "Caesula", "Decima", "Fausta", "Gaia", "Gnaea", "Hosta", "Lucia", "Maio", "Marcia", "Maxima", "Mettia", "Nona", "Numeria", "Octavia", "Postuma", "Prima", "Procula", "Septima", "Servia", "Tertia", "Tiberia", "Titia", "Vibia"},
		Surname: []string{"Antius", "Aurius", "Barbatius", "Calidius", "Cornelius", "Decius", "Fabius", "Flavius", "Galerius", "Horatius", "Julius", "Juventius", "Licinius", "Marius", "Minicius", "Nerius", "Octavius", "Pompeius", "Quinctius", "Rutilius", "Sextius", "Titius", "Ulpius", "Valerius", "Vitellius"},
		Place: []string{"Abilia", "Lucus", "Alsium", "Lugdunum", "Aquileia", "Mediolanum", "Argentoratum", "Novaesium", "Ascrivium", "Patavium", "Asculum", "Pistoria", "Attalia", "Pompeii", "Barium", "Raurica", "Batavorum", "Rigomagus", "Belum", "Roma", "Bobbium", "Salernum", "Brigantium", "Salona", "Burgodunum", "Segovia", "Camulodunum", "Sirmium", "Clausentum", "Spalatum", "Corduba", "Tarraco", "Coriovallum", "Treverorum", "Durucobrivis", "Verulamium", "Eboracum", "Vesontio", "Emona", "Vetera", "Florentia", "Vindelicorum", "Lactodurum", "Vindobona", "Lentia", "Vinovia", "Lindum", "Viroconium", "Londinium", "Volubilis"},
	},
	Nigerian: {
		Male:  []string{"Adesegun", "Akintola", "Amabere", "Arikawe", "Asagwara", "Chidubem", "Chinedu", "Chiwetei", "Damilola", "Esangbedo", "Ezenwoye", "Folarin", "Genechi", "Idowu", "Kelechi", "Ketanndu", "Melubari", "Nkanta", "Obafemi", "Olatunde", "Olumide", "Tombari", "Udofia", "Uyoata", "Uzochi"},
		Female: []string{"Abike", "Adesuwa", "Adunola", "Anguli", "Arewa", "Asari", "Bisola", "Chioma", "Eduwa", "Emilohi", "Fehintola", "Folasade", "Mahparah", "Minika", "Nkolika", "Nkoyo", "Nuanae", "Obioma", "Olafemi", "Shanumi", "Sominabo", "Suliat", "Tariere", "Temedire", "Yemisi"},
		Surname: []string{"Adegboye", "Adeniyi", "Adeyeku", "Adunola", "Agbaje", "Akpan", "Akpehi", "Aliki", "Asuni", "Babangida", "Ekim", "Ezeiruaku", "Fabiola", "Fasola", "Nwokolo", "Nzeocha", "Ojo", "Okonkwo", "Okoye", "Olaniyan", "Olawale", "Olumese", "Onajobi", "Soyinka", "Yamusa"},
		Place: []string{"Abadan", "Jere", "Ador", "Kalabalge", "Agatu", "Katsina", "Akamkpa", "Knoduga", "Akpabuyo", "Konshishatse", "Ala", "Kukawa", "Askira", "Kwande", "Bakassi", "Kwayakusar", "Bama", "Logo", "Bayo", "Mafa", "Bekwara", "Makurdi", "Biase", "Nganzai", "Boki", "Obanliku", "Buruku", "Obi", "Calabar", "Obubra", "Chibok", "Obudu", "Damboa", "Odukpani", "Dikwa", "Ogbadibo", "Etung", "Ohimini", "Gboko", "Okpokwu", "Gubio", "Otukpo", "Guzamala", "Shani", "Gwoza", "Ugep", "Hawul", "Vandeikya", "Ikom", "Yala"},
	},
	Russian: {
		Male:  []string{"Aleksandr", "Andrei", "Arkady", "Boris", "Dmitri", "Dominik", "Grigory", "Igor", "Ilya", "Ivan", "Kiril", "Konstantin", "Leonid", "Nikolai", "Oleg", "Pavel", "Petr", "Sergei", "Stepan", "Valentin", "Vasily", "Viktor", "Yakov", "Yegor", "Yuri"},
		Female: []string{"Aleksandra", "Anastasia", "Anja", "Catarina", "Devora", "Dima", "Ekaterina", "Eva", "Irina", "Karolina", "Katlina", "Kira", "Ludmilla", "Mara", "Nadezdha", "Nastassia", "Natalya", "Oksana", "Olena", "Olga", "Sofia", "Svetlana", "Tatyana", "Vilma", "Yelena"},
		Surname: []string{"Abelev", "Bobrikov", "Chemerkin", "Gogunov", "Gurov", "Iltchenko", "Kavelin", "Komarov", "Korovin", "Kurnikov", "Lebedev", "Litvak", "Mekhdiev", "Muraviov", "Nikitin", "Ortov", "Peshkov", "Romasko", "Shvedov", "Sikorski", "Stolypin", "Turov", "Volokh", "Zaitsev", "Zhukov"},
		Place: []string{"Amur", "Omsk", "Arkhangelsk", "Orenburg", "Astrakhan", "Oryol", "Belgorod", "Penza", "Bryansk", "Perm", "Chelyabinsk", "Pskov", "Chita", "Rostov", "Gorki", "Ryazan", "Irkutsk", "Sakhalin", "Ivanovo", "Samara", "Kaliningrad", "Saratov", "Kaluga", "Smolensk", "Kamchatka", "Sverdlovsk", "Kemerovo", "Tambov", "Kirov", "Tomsk", "Kostroma", "Tula", "Kurgan", "Tver", "Kursk", "Tyumen", "Leningrad", "Ulyanovsk", "Lipetsk", "Vladimir", "Magadan", "Volgograd", "Moscow", "Vologda", "Murmansk", "Voronezh", "Novgorod", "Vyborg", "Novosibirsk", "Yaroslavl"},
	},
	Spanish: {
		Male:  []string{"Alejandro", "Alonso", "Amelio", "Armando", "Bernardo", "Carlos", "Cesar", "Diego", "Emilio", "Estevan", "Felipe", "Francisco", "Guillermo", "Javier", "Jose", "Juan", "Julio", "Luis", "Pedro", "Raul", "Ricardo", "Salvador", "Santiago", "Valeriano", "Vicente"},
		Female: []string{"Adalina", "Aleta", "Ana", "Ascencion", "Beatriz", "Carmela", "Celia", "Dolores", "Elena", "Emelina", "Felipa", "Inez", "Isabel", "Jacinta", "Lucia", "Lupe", "Maria", "Marta", "Nina", "Paloma", "Rafaela", "Soledad", "Teresa", "Valencia", "Zenaida"},
		Surname: []string{"Arellano", "Arispana", "Borrego", "Carderas", "Carranzo", "Cordova", "Enciso", "Espejo", "Gavilan", "Guerra", "Guillen", "Huertas", "Illan", "Jurado", "Moretta", "Motolinia", "Pancorbo", "Paredes", "Quesada", "Roma", "Rubiera", "Santoro", "Torrillas", "Vera", "Vivero"},
		Place: []string{"Aguascebas", "Loreto", "Alcazar", "Lujar", "Barranquete", "Marbela", "Bravatas", "Matagorda", "Cabezudos", "Nacimiento", "Calderon", "Niguelas", "Cantera", "Ogijares", "Castillo", "Ortegicar", "Delgadas", "Pampanico", "Donablanca", "Pelado", "Encinetas", "Quesada", "Estrella", "Quintera", "Faustino", "Riguelo", "Fuentebravia", "Ruescas", "Gafarillos", "Salteras", "Gironda", "Santopitar", "Higueros", "Taberno", "Huelago", "Torres", "Humilladero", "Umbrete", "Illora", "Valdecazorla", "Isabela", "Velez", "Izbor", "Vistahermosa", "Jandilla", "Yeguas", "Jinetes", "Zahora", "Limones", "Zumeta"},
	},
}