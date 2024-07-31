// pages/university/index.jsx
import { useState,useEffect  } from "react";
import Navbar from "../../components/Navbar";
import Swal from "sweetalert2";

import {
  getJadwalKuliah,
  getDosen,
  getMahasiswa,
  createJadwalKuliah,
  deleteJadwalKuliah,
  updateJadwalKuliah,
} from "../../utils/api";

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
    console.error("Error in getServerSideProps:", error);
    return {
      props: {
        jadwalKuliah: [],
        dosen: [],
        mahasiswa: [],
      },
    };
  }
}

const Home = () => {
  const [jadwalKuliah, setJadwalKuliah] = useState([]);
  const [dosen, setDosen] = useState([]);
  const [mahasiswa, setMahasiswa] = useState([]);
  const [form, setForm] = useState({
    dosenId: "",
    mahasiswaId: "",
    hari: "",
    jamMulai: "",
    jamSelesai: "",
  });
  const [editMode, setEditMode] = useState(false);
  const [editId, setEditId] = useState(null);

  // Mengambil data saat komponen dimuat
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [jadwalData, dosenData, mahasiswaData] = await Promise.all([
          getJadwalKuliah(),
          getDosen(),
          getMahasiswa(),
        ]);
        setJadwalKuliah(jadwalData.data);
        setDosen(dosenData.data);
        setMahasiswa(mahasiswaData.data);
      } catch (error) {
        Swal.fire("Error!", "Gagal memuat data.", "error");
        console.error("Failed to fetch data:", error);
      }
    };

    fetchData();
  }, []);
  
  const fetchData = async () => {
    try {
      const response = await getJadwalKuliah();
      if (response && response.status === 200) {
        setJadwalKuliah(response.data || []);
      } else {
        const errorMessage = response ? `Failed to fetch data: ${response.statusText}` : "Failed to fetch data: No response from the server";
        console.error("Failed to fetch data:", response ? response.status : "No status", response ? response.statusText : "No status text");
        Swal.fire("Error!", errorMessage, "error");
      }
    } catch (error) {
      console.error("Error fetching data:", error);
      Swal.fire("Error!", `Exception while fetching data: ${error.toString()}`, "error");
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
  
    const dosenId = parseInt(form.dosenId, 10);
    const mahasiswaId = parseInt(form.mahasiswaId, 10);
  
    const startTime = new Date(`1970-01-01T${form.jamMulai}:00Z`);
    const endTime = new Date(`1970-01-01T${form.jamSelesai}:00Z`);
  
    if (endTime <= startTime) {
      Swal.fire({
        title: "Error!",
        text: "Jam selesai harus lebih besar dari jam mulai!",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }
  
    if (form.jamMulai === form.jamSelesai) {
      Swal.fire({
        title: "Error!",
        text: "Jam mulai dan jam selesai tidak boleh sama!",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }
  
    const existingEntry = jadwalKuliah.find(
      (jk) => jk.dosenId === dosenId &&
              jk.mahasiswaId === mahasiswaId &&
              jk.hari === form.hari &&
              jk.ID !== editId
    );
  
    if (existingEntry) {
      Swal.fire({
        title: "Error!",
        text: "Dosen dan Mahasiswa sudah memiliki jadwal pada hari yang sama!",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }
  
    const submitData = {
      ...form,
      dosenId,
      mahasiswaId,
    };
  
    if (editMode) {
      try {
        await updateJadwalKuliah(editId, submitData);
        Swal.fire("Updated!", "Jadwal Kuliah berhasil diperbarui.", "success");
        setEditMode(false);
        setEditId(null);
        await fetchData();  
      } catch (error) {
        Swal.fire("Error!", error.message, "error");
      }
    } else {
      try {
        await createJadwalKuliah(submitData);
        Swal.fire({
          title: "Berhasil!",
          text: "Jadwal Kuliah berhasil ditambahkan!",
          icon: "success",
          confirmButtonText: "Ok",
        });
        await fetchData();  
      } catch (error) {
        Swal.fire({
          title: "Error!",
          text: error.message,
          icon: "error",
          confirmButtonText: "Ok",
        });
      }
    }
  };
  

  const handleDelete = async (id) => {
    Swal.fire({
      title: "Are you sure?",
      text: "You won't be able to revert this!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Yes, delete it!",
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          await deleteJadwalKuliah(id);
          Swal.fire("Deleted!", "Your file has been deleted.", "success");
          // Optionally refresh the list here
        } catch (error) {
          Swal.fire("Failed!", error.message, "error");
        }
      }
    });
  };

  const handleEdit = (jadwal) => {
    setForm({
      dosenId: jadwal.DosenID.toString(),
      mahasiswaId: jadwal.MahasiswaID.toString(),
      hari: jadwal.Hari,
      jamMulai: formatTimeForInput(jadwal.JamMulai),
      jamSelesai: formatTimeForInput(jadwal.JamSelesai),
    });
    setEditId(jadwal.ID);
    setEditMode(true);
  };

  const formatTimeForInput = (timeString) => {
    return timeString ? timeString.slice(0, 5) : '--:--'; // Slices the string to only show HH:mm
  };
  


  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold">Jadwal Kuliah</h1>
        <form onSubmit={handleSubmit} className="mt-4">
          <div className="mb-4">
            <label className="block text-gray-700">Dosen</label>
            <select
              name="dosenId"
              value={form.dosenId}
              onChange={handleInputChange}
              className="mt-1 block w-full"
            >
              <option value="">Pilih Dosen</option>
              {dosen.map((item) => (
                <option key={item.ID} value={item.ID}>
                  {item.Nama} - {item.MataKuliah.Nama}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Mahasiswa</label>
            <select
              name="mahasiswaId"
              value={form.mahasiswaId}
              onChange={handleInputChange}
              className="mt-1 block w-full"
            >
              <option value="">Pilih Mahasiswa</option>
              {mahasiswa.map((item) => (
                <option key={item.ID} value={item.ID}>
                  {item.Nama}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Hari</label>
            <select
              name="hari"
              value={form.hari}
              onChange={handleInputChange}
              className="mt-1 block w-full"
            >
              <option value="">Pilih Hari</option>
              <option value="Senin">Senin</option>
              <option value="Selasa">Selasa</option>
              <option value="Rabu">Rabu</option>
              <option value="Kamis">Kamis</option>
              <option value="Jumat">Jumat</option>
            </select>
          </div>
          {/* Jam Mulai */}
          <div className="mb-4">
            <label className="block text-gray-700">Jam Mulai</label>
            <input
              type="time"
              name="jamMulai"
              value={form.jamMulai}
              onChange={handleInputChange}
              className="mt-1 block w-full p-2 border border-gray-300 rounded shadow-sm"
            />
          </div>
          {/* Jam Selesai */}
          <div className="mb-4">
            <label className="block text_gray-700">Jam Selesai</label>
            <input
              type="time"
              name="jamSelesai"
              value={form.jamSelesai}
              onChange={handleInputChange}
              className="mt-1 block w-full p-2 border border-gray-300 rounded shadow-sm"
            />
          </div>
          <button
            type="submit"
            className="bg-blue-500 text-white px-4 py-2 rounded"
          >
            Tambah Jadwal
          </button>
        </form>

        <div className="mt-8">
        <h2 className="text-xl font-bold">Daftar Jadwal Kuliah</h2>
        <table className="min-w-full bg-white">
          <thead>
            <tr className="bg-gray-100">
              <th className="py-2">Dosen</th>
              <th className="py-2">Mata Kuliah</th>
              <th className="py-2">Mahasiswa</th>
              <th className="py-2">Hari</th>
              <th className="py-2">Jam Mulai</th>
              <th className="py-2">Jam Selesai</th>
              <th className="py-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {jadwalKuliah.map((jk) => (
              <tr key={jk.ID}>
                <td className="border px-4 py-2">{jk.Dosen.Nama}</td>
                <td className="border px-4 py-2">{jk.Dosen.MataKuliah ? jk.Dosen.MataKuliah.Nama : 'Mata Kuliah tidak tersedia'}</td>
                <td className="border px-4 py-2">{jk.Mahasiswa.Nama}</td>
                <td className="border px-4 py-2">{jk.Hari}</td>
                <td className="border px-4 py-2">{formatTimeForInput(jk.JamMulai)}</td>
                <td className="border px-4 py-2">{formatTimeForInput(jk.JamSelesai)}</td>
                <td className="border px-4 py-2">
                  <button onClick={() => handleEdit(jk)} className="bg-blue-500 text-white px-4 py-2 rounded mr-2">Edit</button>
                  <button onClick={() => handleDelete(jk.ID)} className="bg-red-500 text-white px-4 py-2 rounded">Delete</button>
                </td>
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
