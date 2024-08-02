import { useState, useEffect } from "react";
import Navbar from "../../components/Navbar";
import Swal from "sweetalert2";
import useAuthStore from "../../store/authStore";
import {
  getNilai,
  createNilai,
  updateNilai,
  deleteNilai,
  getMahasiswa,
  getMataKuliah
} from "../../utils/api";

const NilaiPage = () => {
  const { token, user } = useAuthStore();
  const [nilai, setNilai] = useState([]);
  const [mahasiswaList, setMahasiswaList] = useState([]);
  const [mataKuliahList, setMataKuliahList] = useState([]);
  const [form, setForm] = useState({
    indeks: "",
    skor: "",
    mahasiswa_id: "",
    mata_kuliah_id: "",
  });
  const [editMode, setEditMode] = useState(false);
  const [editId, setEditId] = useState(null);

  useEffect(() => {
    fetchNilai();
    fetchMahasiswa();
    fetchMataKuliah();
  }, []);

  const fetchNilai = async () => {
    try {
      const response = await getNilai(token);
      setNilai(response.data);
    } catch (error) {
      Swal.fire("Error", "Login Terlebih Dahulu.", "error");
      console.error("Fetch error:", error);
    }
  };

  const fetchMahasiswa = async () => {
    try {
      const response = await getMahasiswa();
      setMahasiswaList(response.data);
    } catch (error) {
      Swal.fire("Error", "Login Terlebih Dahulu mahasiswa.", "error");
      console.error("Fetch error:", error);
    }
  };

  const fetchMataKuliah = async () => {
    try {
      const response = await getMataKuliah();
      setMataKuliahList(response.data);
    } catch (error) {
      Swal.fire("Error", "Login Terlebih Dahulu mata kuliah.", "error");
      console.error("Fetch error:", error);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      indeks: form.indeks,
      skor: parseInt(form.skor), // Pastikan skor dikirim sebagai integer
      mahasiswa_id: parseInt(form.mahasiswa_id),
      mata_kuliah_id: parseInt(form.mata_kuliah_id),
      users_id: user.id, // Mengirimkan users_id dari user yang sedang login
    };
    try {
      if (editMode) {
        await updateNilai(editId, payload, token);
        Swal.fire("Sukses", "Data nilai berhasil diperbarui!", "success");
        setEditMode(false);
        setEditId(null);
      } else {
        await createNilai(payload, token);
        Swal.fire("Sukses", "Data nilai berhasil ditambahkan!", "success");
      }
      setForm({
        indeks: "",
        skor: "",
        mahasiswa_id: "",
        mata_kuliah_id: "",
      });
      fetchNilai();
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  const handleEdit = (nilai) => {
    setForm({
      indeks: nilai.indeks,
      skor: nilai.skor.toString(),
      mahasiswa_id: nilai.mahasiswa_id.toString(),
      mata_kuliah_id: nilai.mata_kuliah_id.toString(),
    });
    setEditId(nilai.ID);
    setEditMode(true);
  };

  const handleDelete = async (id) => {
    try {
      await deleteNilai(id, token);
      Swal.fire("Deleted!", "Data nilai berhasil dihapus.", "success");
      fetchNilai();
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Kelola Nilai</h1>
        <div className="mb-8">
          <form onSubmit={handleSubmit} className="flex gap-3 mb-4">
            <input
              type="text"
              name="indeks"
              placeholder="Indeks"
              value={form.indeks}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded"
            />
            <input
              type="number"
              name="skor"
              placeholder="Skor"
              value={form.skor}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded"
            />
            <select
              name="mahasiswa_id"
              value={form.mahasiswa_id}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded"
            >
              <option value="">Pilih Mahasiswa</option>
              {mahasiswaList.map((m) => (
                <option key={m.ID} value={m.ID}>
                  {m.Nama}
                </option>
              ))}
            </select>
            <select
              name="mata_kuliah_id"
              value={form.mata_kuliah_id}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded"
            >
              <option value="">Pilih Mata Kuliah</option>
              {mataKuliahList.map((mk) => (
                <option key={mk.ID} value={mk.ID}>
                  {mk.Nama}
                </option>
              ))}
            </select>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
              {editMode ? "Update" : "Tambah"}
            </button>
          </form>
        </div>
        <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Indeks
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Skor
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Mahasiswa
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Mata Kuliah
                </th>
                <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {nilai.map((n) => (
                <tr key={n.ID}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {n.indeks}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {n.skor}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {n.mahasiswa.Nama}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {n.mata_kuliah.Nama}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                    <button
                      onClick={() => handleEdit(n)}
                      className="text-indigo-600 hover:text-indigo-900 px-3 py-1 rounded"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(n.ID)}
                      className="text-red-600 hover:text-red-900 px-3 py-1 rounded"
                    >
                      Delete
                    </button>
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

export default NilaiPage;
