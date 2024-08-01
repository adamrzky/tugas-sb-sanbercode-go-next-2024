// pages/matakuliah/index.jsx
import { useState, useEffect } from "react";
import Navbar from "../../components/Navbar";
import Swal from "sweetalert2";
import {
  getMataKuliah,
  createMataKuliah,
  updateMataKuliah,
  deleteMataKuliah,
} from "../../utils/api";

export async function getServerSideProps() {
  try {
    const response = await getMataKuliah();
    if (response.status = 200) {
        console.log(response);
      return {
        props: { mataKuliah: response.data.data || [] },
      };
    } else {
      console.error("Failed to fetch mata kuliah:", response.statusText);
      return {
        props: { mataKuliah: [] },
      };
    }
  } catch (error) {
    console.error("Error during getServerSideProps:", error);
    return {
      props: { mataKuliah: [] },
    };
  }
}

const MataKuliahPage = ({ mataKuliah }) => {
  const [data, setData] = useState(mataKuliah);
  const [form, setForm] = useState({ id: null, nama: "" });
  const [editMode, setEditMode] = useState(false);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { nama: form.nama };
    try {
      if (editMode) {
        await updateMataKuliah(form.id, payload);
        Swal.fire("Success", "Mata kuliah updated successfully!", "success");
        setEditMode(false);
      } else {
        await createMataKuliah(payload);
        Swal.fire("Success", "Mata kuliah added successfully!", "success");
      }
      fetchMataKuliah();
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
    setForm({ id: null, nama: "" }); // Reset form
  };

  const fetchMataKuliah = async () => {
    const response = await getMataKuliah();
    if (response.status = 200) {
      setData(response.data);
    } else {
      Swal.fire("Error", "Failed to fetch mata kuliah.", "error");
    }
  };

  const handleEdit = (mk) => {
    setForm({ id: mk.ID, nama: mk.Nama });
    setEditMode(true);
  };

  const handleDelete = async (id) => {
    await deleteMataKuliah(id);
    Swal.fire("Deleted!", "Mata kuliah has been deleted.", "success");
    fetchMataKuliah();
  };

  useEffect(() => {
    fetchMataKuliah(); // Refresh data on page load
  }, []);

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Kelola Mata Kuliah</h1>
        <div className="mb-8">
          <form onSubmit={handleSubmit} className="flex gap-3 items-center">
            <input
              type="text"
              name="nama"
              placeholder="Nama Mata Kuliah"
              value={form.nama}
              onChange={handleInputChange}
              required
              className="p-2 border border-gray-300 rounded"
            />
            <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              {editMode ? "Update" : "Tambah"}
            </button>
          </form>
        </div>
        <table className="min-w-full bg-white shadow overflow-hidden rounded-lg">
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
            {data.map((mk) => (
              <tr key={mk.ID}>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {mk.Nama}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                  <button onClick={() => handleEdit(mk)} className="text-indigo-600 hover:text-indigo-900 px-3 py-1 rounded mr-4">
                    Edit
                  </button>
                  <button onClick={() => handleDelete(mk.ID)} className="text-red-600 hover:text-red-900 px-3 py-1 rounded">
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default MataKuliahPage;
