// components/Navbar.jsx
import Link from 'next/link';

const Navbar = () => {
  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-white font-bold text-xl">
          University App
        </div>
        <div className="space-x-4">
          <Link href="/university">
            <span className="text-gray-300 hover:text-white cursor-pointer">Home</span>
          </Link>
          <Link href="/mahasiswa">
            <span className="text-gray-300 hover:text-white cursor-pointer">Mahasiswa</span>
          </Link>
          <Link href="/university/mata-kuliah">
            <span className="text-gray-300 hover:text-white cursor-pointer">Mata Kuliah</span>
          </Link>
          <Link href="/university/dosen">
            <span className="text-gray-300 hover:text-white cursor-pointer">Dosen</span>
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
