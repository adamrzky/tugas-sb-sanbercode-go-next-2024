// utils/api.js
import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

export const getJadwalKuliah = async () => {
  const response = await api.get('/jadwal-kuliah');
  return response.data;
};

export const getDosen = async () => {
  const response = await api.get('/dosen');
  return response.data;
};

export const getMahasiswa = async () => {
  const response = await api.get('/mahasiswa');
  return response.data;
};

export const createJadwalKuliah = async (data) => {
  try {
    const response = await api.post('/jadwal-kuliah', data);
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.error || 'Unknown error occurred');
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


