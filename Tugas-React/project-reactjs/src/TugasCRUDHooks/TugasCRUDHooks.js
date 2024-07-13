import React, { useState } from "react";
import "../App.css";

const dataAwal = [
  { nama: "Nanas", hargaTotal: 100000, beratTotal: 4000 },
  { nama: "Manggis", hargaTotal: 350000, beratTotal: 10000 },
  { nama: "Nangka", hargaTotal: 90000, beratTotal: 2000 },
  { nama: "Durian", hargaTotal: 400000, beratTotal: 5000 },
  { nama: "Strawberry", hargaTotal: 120000, beratTotal: 6000 },
];

function TugasCRUDHooks() {
  const [daftarBuah, setDaftarBuah] = useState(dataAwal);
  const [formData, setFormData] = useState({ nama: "", hargaTotal: "", beratTotal: "" });
  const [isEditing, setIsEditing] = useState(false);
  const [editIndex, setEditIndex] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (formData.nama && formData.hargaTotal && formData.beratTotal) {
      if (isNaN(formData.hargaTotal) || isNaN(formData.beratTotal) || formData.beratTotal < 2000) {
        alert("Harga Total dan Berat Total harus berupa angka, dan Berat Total minimal 2 kg");
        return;
      }

      const buahBaru = {
        nama: formData.nama,
        hargaTotal: parseInt(formData.hargaTotal),
        beratTotal: parseInt(formData.beratTotal),
      };

      if (isEditing) {
        const buahDiperbarui = daftarBuah.map((buah, index) =>
          index === editIndex ? buahBaru : buah
        );
        setDaftarBuah(buahDiperbarui);
        setIsEditing(false);
        setEditIndex(null);
      } else {
        setDaftarBuah([...daftarBuah, buahBaru]);
      }

      setFormData({ nama: "", hargaTotal: "", beratTotal: "" });
    } else {
      alert("Semua inputan wajib diisi");
    }
  };

  const handleEdit = (index) => {
    const buah = daftarBuah[index];
    setFormData(buah);
    setIsEditing(true);
    setEditIndex(index);
  };

  const handleDelete = (index) => {
    const buahDiperbarui = daftarBuah.filter((_, i) => i !== index);
    setDaftarBuah(buahDiperbarui);
  };

  return (
    <div className="TugasCRUDHooks">
      <h1>Daftar Harga Buah</h1>
      <table>
        <thead>
          <tr>
            <th>No</th>
            <th>Nama</th>
            <th>Harga Total</th>
            <th>Berat Total</th>
            <th>Harga per kg</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          {daftarBuah.map((buah, index) => (
            <tr key={index}>
              <td>{index + 1}</td>
              <td>{buah.nama}</td>
              <td>{buah.hargaTotal}</td>
              <td>{buah.beratTotal / 1000} kg</td>
              <td>{buah.hargaTotal / (buah.beratTotal / 1000)}</td>
              <td>
                <button onClick={() => handleEdit(index)}>Edit</button>
                <button onClick={() => handleDelete(index)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <h2>Form Daftar Harga Buah</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Nama: </label>
          <input type="text" name="nama" value={formData.nama} onChange={handleChange} />
        </div>
        <div>
          <label>Harga Total: </label>
          <input type="text" name="hargaTotal" value={formData.hargaTotal} onChange={handleChange} />
        </div>
        <div>
          <label>Berat Total (dalam gram): </label>
          <input type="text" name="beratTotal" value={formData.beratTotal} onChange={handleChange} />
        </div>
        <button type="submit">submit</button>
      </form>
    </div>
  );
}

export default TugasCRUDHooks;
