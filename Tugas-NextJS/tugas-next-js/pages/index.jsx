import Head from 'next/head';
import Link from 'next/link';
import Image from 'next/image';
import shoppingHero from '../public/shoping_hero.svg'; 
export default function Home() {
  return (
    <div className="bg-pink-50 min-h-screen">
      <Head>
        <title>Product Discovery</title>
      </Head>
      
      {/* Navbar */}
      <div className="flex justify-between items-center p-5 bg-white text-gray-700">
        <div>LOGO</div>
        <div className="space-x-4">
          <Link href="/" passHref><span className="cursor-pointer text-black hover:text-pink-500">Home</span></Link>
          <Link href="/products" passHref><span className="cursor-pointer text-black hover:text-pink-500">Product</span></Link>
        </div>
      </div>
      
      {/* Main Banner */}
      <div className="flex justify-around items-center p-10 bg-pink-200">
        <div className="text-center">
          <h1 className="text-4xl font-bold">Temukan produk pilihan kamu disini!</h1>
          <p className="text-md mb-4">Nikmati diskon hingga 100% setiap pembelian yang kamu lakukan</p>
          <button className="px-8 py-3 bg-pink-500 text-white rounded hover:bg-pink-600 transition duration-300">Find Product</button>
        </div>
        <Image src={shoppingHero} alt="Shopping" width={500} height={500}/>
      </div>

      {/* Search Bar */}
      <div className="px-10 py-5">
        <input className="border-2 border-gray-300 rounded p-2 w-full" placeholder="Search Product..." />
      </div>

      {/* Recommended Products */}
      <div className="px-10 py-2 bg-white">
        <h2 className="font-bold text-lg mb-4">This Recommended For You!</h2>
        <div className="flex flex-wrap justify-center gap-5">
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
    </div>
  );
}
