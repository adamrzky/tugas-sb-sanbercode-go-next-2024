import { useState, useEffect } from "react";
import Navbar from "../../components/Navbar";
import Swal from "sweetalert2";
import {
  getMahasiswa,
  createMahasiswa,
  updateMahasiswa,
  deleteMahasiswa,
} from "../../utils/api";

export async function getServerSideProps() {
  try {
    const response = await getMahasiswa();
    if (response.status = 200) {
      return {
        props: {
          data: response.data || [], // Pastikan ini adalah array
        },
      };
    } else {
      console.error("Failed to fetch data:", response.status, response.statusText);
      return {
        props: {
          data: [],
        },
      };
    }
  } catch (error) {
    console.error("Error in getServerSideProps:", error);
    return {
      props: {
        data: [],
      },
    };
  }
}

const MahasiswaPage = ({ data }) => {
  const [mahasiswa, setMahasiswa] = useState(data);
  const [form, setForm] = useState({
    id: null,
    nama: "",
  });
  const [editMode, setEditMode] = useState(false);

  useEffect(() => {
    fetchMahasiswa(); // Memuat ulang data setelah update
  }, []);

  const fetchMahasiswa = async () => {
    const response = await getMahasiswa();
    if (response.status = 200) {
      setMahasiswa(response.data);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { nama: form.nama };

    if (editMode && form.id) {
      try {
        await updateMahasiswa(form.id, payload);
        Swal.fire("Sukses", "Data mahasiswa berhasil diperbarui!", "success");
        setEditMode(false);
        fetchMahasiswa();
      } catch (error) {
        Swal.fire("Error", error.message, "error");
      }
    } else {
      try {
        await createMahasiswa(payload);
        Swal.fire("Sukses", "Mahasiswa berhasil ditambahkan!", "success");
        fetchMahasiswa();
      } catch (error) {
        Swal.fire("Error", error.message, "error");
      }
    }
    setForm({ id: null, nama: "" }); // Reset form
  };

  const handleEdit = (mahasiswa) => {
    setForm({
      id: mahasiswa.ID,
      nama: mahasiswa.Nama,
    });
    setEditMode(true);
  };

  const handleDelete = async (id) => {
    try {
      await deleteMahasiswa(id);
      Swal.fire("Deleted!", "Mahasiswa has been deleted.", "success");
      fetchMahasiswa();
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Kelola Mahasiswa</h1>
        <div className="mb-8">
          <form onSubmit={handleSubmit} className="flex gap-3 mb-4">
            <input
              type="text"
              name="nama"
              placeholder="Nama Mahasiswa"
              value={form.nama}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded shadow-sm"
            />
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
              {editMode ? "Perbarui" : "Tambah"}
            </button>
          </form>
        </div>
        <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Nama
                </th>
                <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {mahasiswa.map((m) => (
                <tr key={m.ID}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {m.Nama}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                    <button
                      onClick={() => handleEdit(m)}
                      className="text-indigo-600 hover:text-indigo-900 px-3 py-1 rounded"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(m.ID)}
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

export default MahasiswaPage;
