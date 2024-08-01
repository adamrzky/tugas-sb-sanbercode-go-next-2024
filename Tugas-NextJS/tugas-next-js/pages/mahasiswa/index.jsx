import { useState } from "react";
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
      const response = await getMahasiswa(); // Sesuaikan ini dengan fungsi yang benar untuk mengambil data mahasiswa
      if (response.status === 200) {
        return {
          props: {
            mahasiswa: response.data.data || [], // Pastikan ini adalah array
          },
        };
      } else {
        console.error("Failed to fetch data:", response.status, response.statusText);
        return {
          props: {
            mahasiswa: [],
          },
        };
      }
    } catch (error) {
      console.error("Error in getServerSideProps:", error);
      return {
        props: {
          mahasiswa: [],
        },
      };
    }
  }
  
  

const MahasiswaPage = ({ mahasiswa }) => {
  const [mahasiswaList, setMahasiswaList] = useState(mahasiswa);
  const [form, setForm] = useState({
    id: null,
    nama: "",
    nim: "",
    jurusan: "",
  });
  const [editMode, setEditMode] = useState(false);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { nama: form.nama, nim: form.nim, jurusan: form.jurusan };

    try {
      let response;
      if (editMode) {
        response = await updateMahasiswa(form.id, payload);
        setEditMode(false);
      } else {
        response = await createMahasiswa(payload);
      }
      if (response.status === 200) {
        setForm({ id: null, nama: "", nim: "", jurusan: "" });
        setMahasiswaList([...mahasiswaList, response.data]); // update list
        Swal.fire("Success", "Mahasiswa updated successfully!", "success");
      } else {
        Swal.fire("Error", response.statusText, "error");
      }
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  const handleEdit = (mahasiswa) => {
    setForm({
      id: mahasiswa.id,
      nama: mahasiswa.nama,
      nim: mahasiswa.nim,
      jurusan: mahasiswa.jurusan,
    });
    setEditMode(true);
  };

  const handleDelete = async (id) => {
    try {
      const response = await deleteMahasiswa(id);
      if (response.status === 200) {
        const newList = mahasiswaList.filter((m) => m.id !== id);
        setMahasiswaList(newList);
        Swal.fire("Deleted!", "Mahasiswa has been deleted.", "success");
      } else {
        Swal.fire("Error", response.statusText, "error");
      }
    } catch (error) {
      Swal.fire("Error", error.message, "error");
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold">Kelola Mahasiswa</h1>
        <form onSubmit={handleSubmit} className="mt-4">
          <input
            type="text"
            name="nama"
            placeholder="Nama Mahasiswa"
            value={form.nama}
            onChange={handleInputChange}
            required
          />
          <input
            type="text"
            name="nim"
            placeholder="NIM Mahasiswa"
            value={form.nim}
            onChange={handleInputChange}
            required
          />
          <input
            type="text"
            name="jurusan"
            placeholder="Jurusan"
            value={form.jurusan}
            onChange={handleInputChange}
            required
          />
          <button type="submit" className="btn-primary">
            {editMode ? "Update" : "Add"}
          </button>
        </form>
        <div>
          <h2 className="text-xl font-bold">Daftar Mahasiswa</h2>
          <table>
            <thead>
              <tr>
                <th>Nama</th>
                <th>NIM</th>
                <th>Jurusan</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {mahasiswaList.map((m) => (
                <tr key={m.id}>
                  <td>{m.nama}</td>
                  <td>{m.nim}</td>
                  <td>{m.jurusan}</td>
                  <td>
                    <button onClick={() => handleEdit(m)}>Edit</button>
                    <button onClick={() => handleDelete(m.id)}>Delete</button>
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
