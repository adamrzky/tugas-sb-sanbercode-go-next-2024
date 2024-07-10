var filterCarsPromise = require('./promise3.js');

console.log('Soal 3');

// Kondisi 1: Mobil berwarna hitam tahun 2019
filterCarsPromise('black', 2019)
  .then(result => {
    console.log(result);
  })
  .catch(error => {
    console.log(error.message);
  });

// Kondisi 2: Mobil berwarna silver tahun 2017
filterCarsPromise('silver', 2017)
  .then(result => {
    console.log(result);
  })
  .catch(error => {
    console.log(error.message);
  });

// Kondisi 3: Mobil berwarna abu-abu tahun 2019 (menggunakan async/await)
(async function() {
  try {
    const result = await filterCarsPromise('grey', 2019);
    console.log(result);
  } catch (error) {
    console.log(error.message);
  }
})();

// Kondisi 4: Mobil berwarna abu-abu tahun 2018 (menggunakan async/await)
(async function() {
  try {
    const result = await filterCarsPromise('grey', 2018);
    console.log(result);
  } catch (error) {
    console.log(error.message);
  }
})();

// Kondisi 5: Mobil berwarna hitam tahun 2020 (menggunakan async/await)
(async function() {
  try {
    const result = await filterCarsPromise('black', 2020);
    console.log(result);
  } catch (error) {
    console.log(error.message);
  }
})();

console.log(); // Baris kosong untuk pemisah
