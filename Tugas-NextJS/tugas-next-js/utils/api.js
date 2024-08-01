// utils/api.js
import axios from "axios";

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

export const getJadwalKuliah = async () => {
  const response = await api.get("/jadwal-kuliah");
  return response.data;
};

export const getDosen = async () => {
  const response = await api.get("/dosen");
  return response.data;
};

export const getMahasiswa = async () => {
  const response = await api.get("/mahasiswa");
  return response.data;
};

export const createJadwalKuliah = async (data) => {
  try {
    const response = await api.post("/jadwal-kuliah", data);
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.error || "Unknown error occurred");
  }
};

export const updateJadwalKuliah = async (id, data) => {
  try {
    const response = await api.put(`/jadwal-kuliah/${id}`, data);
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.error || "Unknown error occurred");
  }
};

export const deleteJadwalKuliah = async (id) => {
  try {
    const response = await api.delete(`/jadwal-kuliah/${id}`);
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.error || "Unknown error occurred");
  }
};

// Mahasiswa

// Membuat data mahasiswa baru
export const createMahasiswa = async (mahasiswa) => {
  const response = await api.post("/mahasiswa", mahasiswa);
  return response.data;
};

// Memperbarui data mahasiswa
export const updateMahasiswa = async (id, mahasiswa) => {
  const response = await api.put(`/mahasiswa/${id}`, mahasiswa);
  return response.data;
};

// Menghapus data mahasiswa
export const deleteMahasiswa = async (id) => {
  const response = await api.delete(`/mahasiswa/${id}`);
  return response.data;
};


export const getMataKuliah = async () => {
  const response = await api.get("/mata-kuliah");
  return response.data;
};



// Fungsi untuk menambahkan Mata Kuliah baru
export async function createMataKuliah(data) {
  console.log(data)
  try {
    const response = await fetch(`${API_URL}/mata-kuliah`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    console.log(response)
    return await response.json();
  } catch (error) {
    console.error("Failed to create mata kuliah:", error);
  }
}

// Fungsi untuk memperbarui Mata Kuliah
export async function updateMataKuliah(id, data) {
  try {
    const response = await fetch(`${API_URL}/mata-kuliah/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    return await response.json();
  } catch (error) {
    console.error("Failed to update mata kuliah:", error);
  }
}

// Fungsi untuk menghapus Mata Kuliah
export async function deleteMataKuliah(id) {
  try {
    const response = await fetch(`${API_URL}/mata-kuliah/${id}`, {
      method: 'DELETE'
    });
    return await response.json();
  } catch (error) {
    console.error("Failed to delete mata kuliah:", error);
  }
}


// Fungsi untuk menambahkan Dosen baru
export async function createDosen(data) {
  try {
    const response = await fetch(`${API_URL}/dosen`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    return await response.json();
  } catch (error) {
    console.error("Failed to create dosen:", error);
  }
}

// Fungsi untuk memperbarui Dosen
export async function updateDosen(id, data) {
  try {
    const response = await fetch(`${API_URL}/dosen/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    return await response.json();
  } catch (error) {
    console.error("Failed to update dosen:", error);
  }
}

// Fungsi untuk menghapus Dosen
export async function deleteDosen(id) {
  try {
    const response = await fetch(`${API_URL}/dosen/${id}`, {
      method: 'DELETE'
    });
    return await response.json();
  } catch (error) {
    console.error("Failed to delete dosen:", error);
  }
}