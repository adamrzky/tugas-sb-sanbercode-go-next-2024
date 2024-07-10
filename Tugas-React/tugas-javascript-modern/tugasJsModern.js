// Soal 1 - Arrow Function

const luasLingkaran = (radius) => {
    const pi = 3.14159;
    return pi * radius * radius;
};

const kelilingLingkaran = (radius) => {
    const pi = 3.14159;
    return 2 * pi * radius;
};

// Contoh penggunaan:
console.log('Soal 1');
console.log(`Luas lingkaran dengan radius 5: ${luasLingkaran(5)}`); // Output: Luas lingkaran dengan radius 5: 78.53975
console.log(`Keliling lingkaran dengan radius 5: ${kelilingLingkaran(5)}`); // Output: Keliling lingkaran dengan radius 5: 31.4159

console.log(); // Baris kosong untuk pemisah


// Soal 2 - Arrow Function dengan Rest Parameter dan Template Literal

const introduce = (...params) => {
    const [name, age, gender, job] = params;
    const title = gender.toLowerCase() === "laki-laki" ? "Pak" : "Bu";
    return `${title} ${name} adalah seorang ${job} yang berusia ${age} tahun`;
};

console.log('Soal 2');

const perkenalanJohn = introduce("john", "30", "Laki-Laki", "penulis");
console.log(perkenalanJohn); // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
 
const perkenalanSarah = introduce("sarah", "28", "Perempuan", "guru");
console.log(perkenalanSarah); // Menampilkan "Bu Sarah adalah seorang guru yang berusia 28 tahun"

console.log(); // Baris kosong untuk pemisah

// Soal 3 - Mengubah kode ES5 menjadi kode ES6

const newFunction = (firstName, lastName) => ({
    firstName,
    lastName,
    fullName() {
        console.log(`${firstName} ${lastName}`);
    }
});

console.log('Soal 3');
console.log(newFunction("John", "Doe").firstName); // Output: John
console.log(newFunction("Richard", "Roe").lastName); // Output: Roe
newFunction("William", "Imoh").fullName(); // Output: William Imoh

console.log(); // Baris kosong untuk pemisah

// Soal 4 - Menggunakan Destructuring

let phone = {
    name: "Galaxy Note 20",
    brand: "Samsung",
    year: 2020,
    colors: ["Mystic Bronze", "Mystic White", "Mystic Black"]
};

const { brand: phoneBrand, name: phoneName, year, colors } = phone;
const [colorBronze, , colorBlack] = colors;

console.log('Soal 4');
console.log(phoneBrand, phoneName, year, colorBlack, colorBronze);
// Output: Samsung Galaxy Note 20 2020 Mystic Black Mystic Bronze

console.log(); // Baris kosong untuk pemisah

// Soal 5 - Menggunakan Spread Operator

let warna = ["biru", "merah", "kuning" , "hijau"];
let dataBukuTambahan = {
    penulis: "john doe",
    tahunTerbit: 2020 
};
let buku = {
    nama: "pemograman dasar",
    jumlahHalaman: 172,
    warnaSampul: ["hitam"]
};

// Tulis kode jawabannya di sini 
buku = {
    ...buku,
    warnaSampul: [...buku.warnaSampul, ...warna],
    ...dataBukuTambahan
};

console.log('Soal 5');
console.log(buku);


console.log(); // Baris kosong untuk pemisah

// Soal 6 - Mengisi data products

const addProducts = (samsung, newProducts) => {
    samsung.products = [...samsung.products, ...newProducts];
    return samsung;
};

let samsung = {
    name: "Samsung",
    products: [
        { name: "Samsung Galaxy Note 10", colors: ["black", "gold", "silver"] },
        { name: "Samsung Galaxy Note 10s", colors: ["blue", "silver"] },
        { name: "Samsung Galaxy Note 20s", colors: ["white", "black"] }
    ]
};

let newProducts = [
    { name: "Samsung Galaxy A52", colors: ["white", "black"] },
    { name: "Samsung Galaxy M52", colors: ["blue", "grey", "white"] }
];

samsung = addProducts(samsung, newProducts);

console.log('Soal 6');
// console.log(samsung); 
console.log(JSON.stringify(samsung, null, 2)); // Karna running di console

console.log(); // Baris kosong untuk pemisah


// Soal 7 - Konversi Object

const konversiObject = (nama, domisili, umur) => {
    return { nama, domisili, umur };
};

let data = ["Bondra", "Medan", 25];
const [nama, domisili, umur] = data;

console.log('Soal 7');
console.log(konversiObject(nama, domisili, umur));
// Output:
// { "nama" : "Bondra", "domisili": "Medan", "umur": 25 }

console.log(); // Baris kosong untuk pemisah



// Soal 8 - Graduate

const graduate = (...students) => {
    return students.reduce((acc, student) => {
        student.forEach(({ name, class: cls }) => {
            if (!acc[cls]) {
                acc[cls] = [];
            }
            acc[cls].push(name);
        });
        return acc;
    }, {});
};

// TEST CASES
const data1 = [
    { name: "Ahmad", class: "adonis" },
    { name: "Regi", class: "laravel" },
    { name: "Bondra", class: "adonis" },
    { name: "Iqbal", class: "vuejs" },
    { name: "Putri", class: "laravel" }
];

const data2 = [
    { name: "Yogi", class: "react" },
    { name: "Fikri", class: "agile" },
    { name: "Arief", class: "agile" }
];

console.log('Soal 8');
console.log(graduate(data1));


console.log(graduate(data2));


console.log(); // Baris kosong untuk pemisah