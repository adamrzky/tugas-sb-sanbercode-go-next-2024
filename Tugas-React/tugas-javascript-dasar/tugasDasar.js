// Soal 1
function introduce(nama, gender, pekerjaan, umur) {
    let panggilan = gender === "laki-laki" ? "Pak" : "Bu";
    return `${panggilan} ${nama} adalah seorang ${pekerjaan} yang berusia ${umur} tahun`;
}

console.log('Soal 1');
var john = introduce("John", "laki-laki", "penulis", "30");
console.log(john); // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

var sarah = introduce("Sarah", "perempuan", "model", "28");
console.log(sarah); // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

console.log(); 

// Soal 2
function findUniqueCharacters(text) {
    let charCount = {};
    for (let char of text.toLowerCase().replace(/\s/g, '')) {
        charCount[char] = (charCount[char] || 0) + 1;
    }
    let uniqueChars = '';
    for (let char in charCount) {
        if (charCount[char] === 1) {
            uniqueChars += char;
        }
    }
    return uniqueChars;
}

console.log('Soal 2');
var text = "Super Bootcamp Fullstack Dev 2022";
console.log(findUniqueCharacters(text)); // Menampilkan "rbmfkdv0"

console.log(); 

// Soal 3
function findMaxMin(angka) {
    let max = Math.max(...angka);
    let min = Math.min(...angka);
    return `angka terbesar adalah ${max} dan angka terkecil adalah ${min}`;
}

console.log('Soal 3');
var angka = [2, 3, 1, 9, 12, 8, 9, 7];
console.log(findMaxMin(angka)); // Menampilkan "angka terbesar adalah 12 dan angka terkecil adalah 1"

console.log(); 


// Soal 4
function findLongestName(names) {
    let longestName = "";
    for (let name of names) {
        if (name.length > longestName.length) {
            longestName = name;
        }
    }
    return longestName;
}

console.log('Soal 4');
var names = ["Andrew Gillett", "Chris Sawyer", "David Walsh", "John D Rockefeller"];
console.log(findLongestName(names)); // Menampilkan "John D Rockefeller"

console.log(); 


// Soal 5
function arrangeString(str) {
    return str.split('').sort().join('');
}

console.log('Soal 5');
// TEST CASE 
console.log(arrangeString("bahasa")); // Output: aaabhs
console.log(arrangeString("similikiti")); // Output: iiiiiklmst
console.log(arrangeString("sanbercode")); // Output: abcdeenors
console.log(arrangeString("")); // Output: ""

console.log(); 

// Soal 6
function compressString(str) {
    // Susun string secara alfabetis
    let sortedStr = str.split('').sort().join('');

    let compressedStr = '';
    let currentChar = sortedStr[0];
    let count = 0;

    // Loop untuk menghitung karakter dan membuat string terkompresi
    for (let char of sortedStr) {
        if (char === currentChar) {
            count++;
        } else {
            compressedStr += currentChar + count;
            currentChar = char;
            count = 1;
        }
    }
    // Tambahkan karakter terakhir dan hitungannya
    compressedStr += currentChar + count;

    // Kembalikan string terkompresi jika lebih pendek dari string awal
    return compressedStr.length < str.length ? compressedStr : sortedStr;
}

console.log('Soal 6');
// TEST CASES
console.log(compressString("abrakadabra")); // Output: a5b2d1k1r2
console.log(compressString("aabcccccaaa")); // Output: a5b1c5
console.log(compressString("abdul")); // Output: abdlu
console.log(compressString("maman")); // Output: aamn

console.log(); 

// Soal 7

// Soal Palindrome
function palindrome(kata) {
    // Hapus spasi dan ubah ke huruf kecil untuk konsistensi
    let cleanedStr = kata.replace(/\s+/g, '').toLowerCase();
    // Balik string yang sudah dibersihkan
    let reversedStr = cleanedStr.split('').reverse().join('');
    // Periksa apakah string yang dibalik sama dengan string asli
    return cleanedStr === reversedStr;
}

console.log('Soal 7');
// TEST CASES
console.log(palindrome('katak')); // true
console.log(palindrome('blanket')); // false
console.log(palindrome('nababan')); // true
console.log(palindrome('haji ijah')); // true
console.log(palindrome('mister')); // false

console.log(); 

// Soal 8
function isPalindrome(num) {
    let str = num.toString();
    let reversedStr = str.split('').reverse().join('');
    return str === reversedStr;
}

function angkaPalindrome(num) {
    num++;
    while (!isPalindrome(num)) {
        num++;
    }
    return num;
}

console.log('Soal 8');
// TEST CASES
console.log(angkaPalindrome(8)); // 9
console.log(angkaPalindrome(10)); // 11
console.log(angkaPalindrome(117)); // 121
console.log(angkaPalindrome(175)); // 181
console.log(angkaPalindrome(1000)); // 1001

console.log(); 

// Soal 9
function pasanganTerbesar(num) {
    let str = num.toString();
    let maxPair = 0;

    for (let i = 0; i < str.length - 1; i++) {
        let pair = parseInt(str.substring(i, i + 2));
        if (pair > maxPair) {
            maxPair = pair;
        }
    }
    
    return maxPair;
}

console.log('Soal 9');
// TEST CASES
console.log(pasanganTerbesar(641573)); // 73
console.log(pasanganTerbesar(12783456)); // 83
console.log(pasanganTerbesar(910233)); // 91
console.log(pasanganTerbesar(71856421)); // 85
console.log(pasanganTerbesar(79918293)); // 99

console.log(); 

// Soal 10
function cekPermutasi(str1, str2) {
    // Hilangkan spasi dan ubah ke huruf kecil untuk konsistensi
    let cleanedStr1 = str1.replace(/\s+/g, '').toLowerCase();
    let cleanedStr2 = str2.replace(/\s+/g, '').toLowerCase();

    // Jika panjang string berbeda, maka bukan permutasi
    if (cleanedStr1.length !== cleanedStr2.length) {
        return false;
    }

    // Ubah string menjadi array karakter, urutkan, dan gabungkan kembali menjadi string
    let sortedStr1 = cleanedStr1.split('').sort().join('');
    let sortedStr2 = cleanedStr2.split('').sort().join('');

    // Periksa apakah kedua string yang sudah diurutkan sama
    return sortedStr1 === sortedStr2;
}

console.log('Soal 10');
// TEST CASES
console.log(cekPermutasi("abah", "baha")); // true
console.log(cekPermutasi("ondel", "delon")); // true
console.log(cekPermutasi("paul sernine", "arsene lupin")); // true
console.log(cekPermutasi("taco", "taca")); // false

console.log(); 

// Soal 11
function urlify(str, length) {
    let urlifiedStr = '';
    for (let i = 0; i < length; i++) {
        if (str[i] === ' ') {
            urlifiedStr += '%20';
        } else {
            urlifiedStr += str[i];
        }
    }
    return urlifiedStr;
}

console.log('Soal 11');
// TEST CASES
console.log(urlify("Mr John Smith    ", 13)); // Mr%20John%20Smith
console.log(urlify("Bizzare world of Javascript     ", 27)); // Bizzare%20world%20of%20Javascript

console.log(); 

// Soal 12
var arrayDaftarPeserta = ["John Doe", "laki-laki", "baca buku", 1992];
var objDaftarPeserta = {
    nama: arrayDaftarPeserta[0],
    jenisKelamin: arrayDaftarPeserta[1],
    hobi: arrayDaftarPeserta[2],
    tahunLahir: arrayDaftarPeserta[3]
};

console.log('Soal 12');
console.log(objDaftarPeserta);
// Output: { nama: 'John Doe', jenisKelamin: 'laki-laki', hobi: 'baca buku', tahunLahir: 1992 }

console.log(); 


// Soal 13
var sentence = "Super Bootcamp Golang Nextjs 2024";
var result = '';
var vowels = 'aeiouAEIOU0123456789';

for (let i = 0; i < sentence.length; i++) {
    if (vowels.indexOf(sentence[i]) !== -1) {
        result += sentence[i];
    }
}

console.log('Soal 13');
console.log(result);
// Output: "ueoocaOaextjs2024"

console.log(); 


// Soal 14
var fruits = [
    {
        nama: "Nanas",
        warna: "Kuning",
        adaBijinya: "tidak",
        harga: 9000
    },
    {
        nama: "Jeruk",
        warna: "Oranye",
        adaBijinya: "ada",
        harga: 8000
    },
    {
        nama: "Semangka",
        warna: "Hijau & Merah",
        adaBijinya: "ada",
        harga: 10000
    },
    {
        nama: "Pisang",
        warna: "Kuning",
        adaBijinya: "tidak",
        harga: 5000
    }
];

var fruitsWithoutSeeds = fruits.filter(function(fruit) {
    return fruit.adaBijinya === "tidak";
});

console.log('Soal 14');
console.log(fruitsWithoutSeeds);
// Output: [ { nama: 'Nanas', warna: 'Kuning', adaBijinya: 'tidak', harga: 9000 }, { nama: 'Pisang', warna: 'Kuning', adaBijinya: 'tidak', harga: 5000 } ]

console.log(); 



// Soal 15
var people = [
    { name: "John", job: "Programmer", gender: "male", age: 30 },
    { name: "Sarah", job: "Model", gender: "female", age: 27 },
    { name: "Jack", job: "Engineer", gender: "male", age: 25 },
    { name: "Ellie", job: "Designer", gender: "female", age: 35 },
    { name: "Danny", job: "Footballer", gender: "male", age: 30 },
];

var malesAbove29 = people.filter(function(person) {
    return person.gender === "male" && person.age > 29;
});

console.log('Soal 15');
console.log(malesAbove29);
// Output: [ { name: 'John', job: 'Programmer', gender: 'male', age: 30 }, { name: 'Danny', job: 'Footballer', gender: 'male', age: 30 } ]

console.log(); 

// Soal 16
var people = [
    { name: "John", job: "Programmer", gender: "male", age: 30 },
    { name: "Sarah", job: "Model", gender: "female", age: 27 },
    { name: "Jack", job: "Engineer", gender: "male", age: 25 },
    { name: "Ellie", job: "Designer", gender: "female", age: 35 },
    { name: "Danny", job: "Footballer", gender: "male", age: 30 }
];

function getAverageAge(people) {
    var totalAge = 0;
    for (var i = 0; i < people.length; i++) {
        totalAge += people[i].age;
    }
    return totalAge / people.length;
}

console.log('Soal 16');
console.log(getAverageAge(people)); // Output: 29.4

console.log(); 

// Soal 17
var people = [
    { name: "John", job: "Programmer", gender: "male", age: 30 },
    { name: "Sarah", job: "Model", gender: "female", age: 27 },
    { name: "Jack", job: "Engineer", gender: "male", age: 25 },
    { name: "Ellie", job: "Designer", gender: "female", age: 35 },
    { name: "Danny", job: "Footballer", gender: "male", age: 30 }
];

function sortPeopleByAge(people) {
    return people.sort(function(a, b) {
        return a.age - b.age;
    });
}

console.log('Soal 17');
var sortedPeople = sortPeopleByAge(people);
for (var i = 0; i < sortedPeople.length; i++) {
    console.log((i + 1) + '. ' + sortedPeople[i].name);
}
// Output:
// 1. Jack
// 2. Sarah
// 3. Danny
// 4. John
// 5. Ellie

console.log(); 

// Soal 18
var phone = {
    name: "Samsung Galaxy Note 20",
    brand: "Samsung",
    colors: ["Black"],
    release: 2020
};

function addColors(color) {
    phone.colors.push(color);
}

addColors("Gold");
addColors("Silver");
addColors("Brown");

console.log('Soal 18');
console.log(phone);
// Output:
// {
//   name: 'Samsung Galaxy Note 20',
//   brand: 'Samsung',
//   colors: [ 'Black', 'Gold', 'Silver', 'Brown' ],
//   release: 2020
// }

console.log(); 

// Soal 19
var phones = [
    { name: "Samsung Galaxy A52", brand: "Samsung", year: 2021, colors: ["black", "white"] },
    { name: "Redmi Note 10 Pro", brand: "Xiaomi", year: 2021, colors: ["white", "blue"] },
    { name: "Redmi Note 9 Pro", brand: "Xiaomi", year: 2020, colors: ["white", "blue", "black"] },
    { name: "Iphone 12", brand: "Apple", year: 2020, colors: ["silver", "gold"] },
    { name: "Iphone 11", brand: "Apple", year: 2019, colors: ["gold", "black", "silver"] },
];

function filterAndSortPhones(phones) {
    var filteredPhones = phones.filter(function(phone) {
        return phone.colors.includes("black");
    });

    filteredPhones.sort(function(a, b) {
        return a.year - b.year;
    });

    return filteredPhones;
}

function displayPhoneDetails(phones, index = 0) {
    if (index >= phones.length) return;

    console.log(`${index + 1}. ${phones[index].name}, colors available : ${phones[index].colors.join(", ")}`);
    displayPhoneDetails(phones, index + 1);
}

console.log('Soal 19');
var filteredPhones = filterAndSortPhones(phones);
displayPhoneDetails(filteredPhones);
// Output:
// 1. Iphone 11, colors available : gold, black, silver
// 2. Redmi Note 9 Pro, colors available : white, blue, black
// 3. Samsung Galaxy A52, colors available : black, white

console.log(); 

