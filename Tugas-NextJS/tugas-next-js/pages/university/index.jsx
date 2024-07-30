// pages/university/index.jsx
import { useState } from 'react';
import Navbar from '../../components/Navbar';
import { getJadwalKuliah, getDosen, getMahasiswa, createJadwalKuliah, deleteJadwalKuliah } from '../../utils/api';

export async function getServerSideProps() {
  try {
    const jadwalKuliahRes = await getJadwalKuliah();
    const dosenRes = await getDosen();
    const mahasiswaRes = await getMahasiswa();

    return {
      props: {
        jadwalKuliah: jadwalKuliahRes.data || [],
        dosen: dosenRes.data || [],
        mahasiswa: mahasiswaRes.data || [],
      },
    };
  } catch (error) {
    console.error('Error in getServerSideProps:', error);
    return {
      props: {
        jadwalKuliah: [],
        dosen: [],
        mahasiswa: [],
      },
    };
  }
}

const Home = ({ jadwalKuliah, dosen, mahasiswa }) => {
  const [form, setForm] = useState({
    dosenId: '',
    mahasiswaId: '',
    hari: '',
    jamMulai: '',
    jamSelesai: '',
  });


  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm({
      ...form,
      [name]: value,
    });
  };


  const handleSubmit = async (e) => {
    e.preventDefault();

    // Konversi DosenID dan MahasiswaID dari string ke integer
    const dosenId = parseInt(form.dosenId, 10);
    const mahasiswaId = parseInt(form.mahasiswaId, 10);

    // Parsing waktu untuk memastikan validasi waktu lebih akurat
    const startTime = new Date(`1970-01-01T${form.jamMulai}:00Z`);
    const endTime = new Date(`1970-01-01T${form.jamSelesai}:00Z`);

    // Validasi waktu selesai harus lebih besar dari waktu mulai
    if (endTime <= startTime) {
      alert('Jam selesai harus lebih besar dari jam mulai!');
      return;
    }

    // Validasi waktu mulai dan selesai tidak boleh sama
    if (form.jamMulai === form.jamSelesai) {
      alert('Jam mulai dan jam selesai tidak boleh sama!');
      return;
    }

    // Cek entri yang ada untuk menghindari jadwal bentrok
    const existingEntry = jadwalKuliah.find(jk => jk.dosenId === dosenId && jk.mahasiswaId === mahasiswaId && jk.hari === form.hari);
    if (existingEntry) {
      alert('Dosen dan Mahasiswa sudah memiliki jadwal pada hari yang sama!');
      return;
    }

    // Data yang akan dikirim ke backend, memastikan kunci yang dikirim sesuai
    const submitData = {
        ...form,
        dosenId,      // Pastikan nama kunci sesuai dengan yang di backend
        mahasiswaId  // Pastikan nama kunci sesuai dengan yang di backend
    };

    try {
      const result = await createJadwalKuliah(submitData);
      alert('Jadwal Kuliah berhasil ditambahkan!');
      // Opsional: muat ulang data di sini jika diperlukan, misalnya dengan memanggil fungsi untuk memuat jadwal kuliah lagi
    } catch (error) {
      alert('Error: ' + error.message);
    }
};


  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold">Jadwal Kuliah</h1>
        <form onSubmit={handleSubmit} className="mt-4">
          <div className="mb-4">
            <label className="block text-gray-700">Dosen</label>
             <select name="dosenId" value={form.dosenId} onChange={handleInputChange} className="mt-1 block w-full">
              <option value="">Pilih Dosen</option>
              {dosen.map((item) => (
                <option key={item.ID} value={item.ID}>{item.Nama} - {item.MataKuliah.Nama}</option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Mahasiswa</label>
            <select name="mahasiswaId" value={form.mahasiswaId} onChange={handleInputChange} className="mt-1 block w-full">
              <option value="">Pilih Mahasiswa</option>
              {mahasiswa.map((item) => (
                <option key={item.ID} value={item.ID}>{item.Nama}</option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Hari</label>
            <select name="hari" value={form.hari} onChange={handleInputChange} className="mt-1 block w-full">
              <option value="">Pilih Hari</option>
              <option value="Senin">Senin</option>
              <option value="Selasa">Selasa</option>
              <option value="Rabu">Rabu</option>
              <option value="Kamis">Kamis</option>
              <option value="Jumat">Jumat</option>
            </select>
          </div>
          <div className="mb-4">
            <label className="block text_gray-700">Jam Mulai</label>
            <input type="time" name="jamMulai" value={form.jamMulai} onChange={handleInputChange} className="mt-1 block w-full" />
          </div>
          <div className="mb-4">
            <label className="block text_gray-700">Jam Selesai</label>
            <input type="time" name="jamSelesai" value={form.jamSelesai} onChange={handleInputChange} className="mt-1 block w-full" />
          </div>
          <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">Tambah Jadwal</button>
        </form>

        <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold">Jadwal Kuliah</h1>
        <table className="min-w-full bg-white mt-4">
          <thead>
            <tr>
              <th className="border px-4 py-2">Dosen</th>
              <th className="border px-4 py-2">Mata Kuliah</th>
              <th className="border px-4 py-2">Mahasiswa</th>
              <th className="border px-4 py-2">Hari</th>
              <th className="border px-4 py-2">Jam Mulai</th>
              <th className="border px-4 py-2">Jam Selesai</th>
            </tr>
          </thead>
          <tbody>
            {jadwalKuliah.map((jk) => (
              <tr key={jk.ID}>
                <td className="border px-4 py-2">{jk.Dosen.Nama}</td>
                <td className="border px-4 py-2">{jk.Dosen.MataKuliah && jk.Dosen.MataKuliah.Nama ? jk.Dosen.MataKuliah.Nama : "Mata Kuliah tidak tersedia"}</td>
                <td className="border px-4 py-2">{jk.Mahasiswa.Nama}</td>
                <td className="border px-4 py-2">{jk.Hari}</td>
                <td className="border px-4 py-2">{jk.JamMulai}</td>
                <td className="border px-4 py-2">{jk.JamSelesai}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      </div>
    </>
  );
};

export default Home;
