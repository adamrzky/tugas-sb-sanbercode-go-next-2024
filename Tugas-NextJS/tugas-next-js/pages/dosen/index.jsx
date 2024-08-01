// pages/dosen/index.jsx
import { useState, useEffect } from "react";
import Navbar from "../../components/Navbar";
import Swal from "sweetalert2";
import {
  getDosen,
  createDosen,
  updateDosen,
  deleteDosen,
} from "../../utils/api";

export async function getServerSideProps() {
  try {
    const response = await getDosen();
    if (response.status = 200) {
      return {
        props: {
          data: response.data.data || [], // Pastikan ini adalah array
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

const DosenPage = ({ data }) => {
  const [dosen, setDosen] = useState(data);
  const [form, setForm] = useState({
    id: null,
    nama: "",
  });

  const [editMode, setEditMode] = useState(false);

  useEffect(() => {
    fetchDosen();
  }, []);

  const fetchDosen = async () => {
    const response = await getDosen();
    if (response.status = 200) {
      setDosen(response.data);
    } else {
      Swal.fire("Error", "Failed to fetch dosen.", "error");
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { nama: form.nama };

    if (editMode && form.id) {
      try {
        await updateDosen(form.id, payload);
        Swal.fire("Success", "Dosen updated successfully!", "success");
        setEditMode(false);
        fetchDosen();
      } catch (error) {
        Swal.fire("Error", error.message, "error");
      }
    } else {
      try {
        await createDosen(payload);
        Swal.fire("Success", "Dosen added successfully!", "success");
        fetchDosen();
      } catch (error) {
        Swal.fire("Error", error.message, "error");
      }
    }
    setForm({ id: null, nama: "" }); // Reset form
  };

  const handleEdit = (dosen) => {
    setForm({
      id: dosen.ID,
      nama: dosen.Nama,
    });
    setEditMode(true);
  };

  const handleDelete = async (id) => {
    try {
      await deleteDosen(id);
      Swal.fire("Deleted!", "Dosen has been deleted.", "success");
      fetchDosen();
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Kelola Dosen</h1>
        <div className="mb-8">
          <form onSubmit={handleSubmit} className="flex gap-3 mb-4">
            <input
              type="text"
              name="nama"
              placeholder="Nama Dosen"
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
              {dosen.map((d) => (
                <tr key={d.ID}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {d.Nama}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                    <button
                      onClick={() => handleEdit(d)}
                      className="text-indigo-600 hover:text-indigo-900 px-3 py-1 rounded"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(d.ID)}
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

export default DosenPage;
