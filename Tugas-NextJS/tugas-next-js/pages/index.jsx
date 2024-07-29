import Head from 'next/head';
import Link from 'next/link';
import Image from 'next/image';
import { useState } from 'react';
import shoppingHero from '../public/shoping_hero.svg';

export default function Home() {
  const [isMenuOpen, setMenuOpen] = useState(false);

  return (
    <div className="bg-pink-50 min-h-screen">
      <Head>
        <title>Product Discovery</title>
      </Head>
      
      {/* Navbar */}
      <div className="flex justify-between items-center px-5 py-3 bg-white text-gray-700">
        <div>LOGO</div>
        <div onClick={() => setMenuOpen(!isMenuOpen)} className="md:hidden">
          {/* Menu Hamburger Icon */}
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </div>
        <div className={`space-x-4 ${isMenuOpen ? 'flex' : 'hidden'} flex-col md:flex-row md:block`}>
          <Link href="/" passHref><span className="cursor-pointer text-black hover:text-pink-500">Home</span></Link>
          <Link href="/products" passHref><span className="cursor-pointer text-black hover:text-pink-500">Product</span></Link>
        </div>
      </div>
      
      {/* Main Banner */}
      <div className="flex flex-col md:flex-row justify-around items-center py-10 px-4 bg-pink-200">
        <div className="text-center mb-4 md:mb-0">
          <h1 className="text-4xl font-bold">Temukan produk pilihan kamu disini!</h1>
          <p className="text-md mb-4">Nikmati diskon hingga 100% setiap pembelian yang kamu lakukan</p>
          <button className="bg-pink-500 hover:bg-pink-600 text-white font-bold py-2 px-4 rounded">Find Product</button>
        </div>
        <Image src={shoppingHero} alt="Shopping" width={500} height={500} />
      </div>

      {/* Search and Category Links */}
      <div className="bg-white py-4 px-5 flex flex-col md:flex-row items-center justify-between">
      <div className="flex items-center space-x-2 w-full mb-4 md:mb-0">
        <div className="bg-pink-200 p-2 rounded-full">
          <svg className="w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <input className="flex-grow p-2 border border-gray-300 rounded" placeholder="Search Product..." />
      </div>
        <div className="flex items-center space-x-4 text-sm font-semibold text-gray-700">
          <a className="hover:text-pink-500 transition duration-300">Shirt</a>
          <a className="hover:text-pink-500 transition duration-300">Elektronik</a>
          <a className="hover:text-pink-500 transition duration-300">Game</a>
          <a className="hover:text-pink-500 transition duration-300">Hijab</a>
          <a className="hover:text-pink-500 transition duration-300">Shoes</a>
          <a className="hover:text-pink-500 transition duration-300">Laptop</a>
        </div>
      </div>
      {/* Recommended Products */}
      <div className="px-10 py-2 bg-white">
        <h2 className="font-bold text-lg mb-4">This Recommended For You!</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-5">
          <div className="relative border border-gray-100" style={{ width: "300px" }}>
            <div className="relative object-cover w-full">
              <img src="https://fitinline.com/data/article/20210909/Foto-Produk-Baju-001.jpg" alt="Product" />
            </div>
            <div className="p-6">
              <small>
                <span className="bg-green-100 text-green-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">Kategori Produk</span>
              </small>
              <h5 className="mt-4">Nama Produk</h5>
              <ul className="mt-5 text-sm font-thin text-gray-500">
                <li>Stock : stok produk</li>
                <li className="text-lg font-bold">Harga : Rp Produk Harga</li>
              </ul>
              <div className="flex items-center justify-between mt-4 border">
                <button className="h-full px-2 text-black bg-gray-200">-</button>
                <input className="inline-block w-full h-full text-center focus:outline-none" placeholder="1" />
                <button className="h-full px-2 text-black bg-gray-200">+</button>
              </div>
              <button className="block w-full p-4 mt-5 text-sm font-medium text-white bg-rose-400 border rounded-sm" type="button">
                Add to Cart
              </button>
              <Link href="/product-detail" passHref>
                <span className="block w-full p-4 mt-2 text-sm font-medium text-center text-rose-400 bg-white border border-rose-400 rounded-sm cursor-pointer">
                  Detail Product
                </span>
              </Link>
            </div>
          </div>
        </div>
      </div>
      <div className="bg-white text-center p-4 mt-8">
        <p className="text-gray-700 text-sm">
          Created by <span className="font-semibold">Jhon</span> 
          <svg className="inline w-4 h-4 text-red-500 fill-current" viewBox="0 0 24 24">
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
          </svg>
          , Student Sanbercode Batch 40
        </p>
      </div>
    </div>
  );
}
