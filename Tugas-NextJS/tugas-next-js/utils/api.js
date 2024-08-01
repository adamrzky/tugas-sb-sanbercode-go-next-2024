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
