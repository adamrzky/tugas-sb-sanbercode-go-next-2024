import Link from 'next/link';
import useAuthStore from '../store/authStore'; // Pastikan path ini sesuai dengan lokasi authStore Anda
import { useEffect, useState } from 'react';

const Navbar = () => {
  const { user, login, logout } = useAuthStore();
  const [clientUser, setClientUser] = useState(null);

  useEffect(() => {
    setClientUser(user);
  }, [user]);

  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-white font-bold text-xl">University App</div>
        <div className="space-x-4 flex items-center">
          <Link href="/university" legacyBehavior><a className="text-gray-300 hover:text-white">Home</a></Link>
          <Link href="/mahasiswa" legacyBehavior><a className="text-gray-300 hover:text-white">Mahasiswa</a></Link>
          <Link href="/matakuliah" legacyBehavior><a className="text-gray-300 hover:text-white">Mata Kuliah</a></Link>
          <Link href="/dosen" legacyBehavior><a className="text-gray-300 hover:text-white">Dosen</a></Link>
          <Link href="/nilai" legacyBehavior><a className="text-gray-300 hover:text-white">Nilai</a></Link>
          
          {clientUser ? (
            <>
              <span className="text-white mx-2">{clientUser.username}</span>
              <button onClick={logout} className="text-gray-300 hover:text-white">Logout</button>
            </>
          ) : (
            <>
              <Link href="/login" legacyBehavior><a className="text-gray-300 hover:text-white px-2">Login</a></Link>
              <Link href="/register" legacyBehavior><a className="text-gray-300 hover:text-white px-2">Daftar</a></Link>
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
