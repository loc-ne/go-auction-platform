"use client";

import Image from "next/image";
import { useState } from "react";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import LoginModal from "../components/LoginModal";
import RegisterModal from "../components/RegisterModal";

export default function Home() {
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const [isRegisterModalOpen, setIsRegisterModalOpen] = useState(false);

  const featuredItems = [
    {
      id: 1,
      title: "Air Jordan 1 Retro High 'Chicago'",
      description: "Size 42, Fullbox. Đế giày còn cực mới, da mềm. Được săn đón nhất!",
      currentBid: "8.500.000 ₫",
      timeRemaining: "Còn 02:15:30",
      image: "/limited_sneaker.png",
      bids: 42,
      isLive: true,
      seller: "KicksHunter",
    },
    {
      id: 2,
      title: "Rolex Submariner Date 116610LN",
      description: "Đồng hồ cơ Thụy Sỹ đẳng cấp. Tình trạng xuất sắc, kèm giấy tờ khai sinh Zin.",
      currentBid: "215.000.000 ₫",
      timeRemaining: "Còn 12:10:00",
      image: "/retro_camera.png",
      bids: 56,
      isLive: true,
      seller: "LuxuryWatch.vn",
    },
    {
      id: 3,
      title: "Canon AE-1 Vintage + Lens 50mm",
      description: "Dòng máy ảnh film cổ điển huyền thoại. Tình trạng hoạt động hoàn hảo.",
      currentBid: "2.800.000 ₫",
      timeRemaining: "Còn 05:25:12",
      image: "/retro_camera.png",
      bids: 18,
      isLive: false,
      seller: "VintageVibes",
    },
  ];

  return (
    <div className="min-h-screen bg-white dark:bg-[#0f1115] text-gray-900 dark:text-[#e2e8f0] font-sans selection:bg-blue-500 selection:text-white flex flex-col pt-20">
      <Navbar onLoginClick={() => setIsLoginModalOpen(true)} onRegisterClick={() => setIsRegisterModalOpen(true)} />
      <main className="relative py-20 lg:py-32 overflow-hidden flex-1">
        <div className="hidden dark:block absolute top-0 right-1/4 w-96 h-96 bg-blue-600/20 rounded-full blur-[100px] pointer-events-none mix-blend-screen"></div>
        <div className="hidden dark:block absolute bottom-1/4 left-1/4 w-80 h-80 bg-indigo-600/20 rounded-full blur-[100px] pointer-events-none mix-blend-screen"></div>

        <div className="container relative z-10 px-6 mx-auto flex flex-col lg:flex-row items-center gap-12 lg:gap-20">
          <div className="flex-1 text-center lg:text-left">
            <h1 className="text-5xl md:text-6xl lg:text-7xl font-extrabold tracking-tight mb-6 leading-tight text-gray-900 dark:text-white drop-shadow-sm">
              Săn đồ hiếm, <br className="hidden md:block" />
              <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 dark:from-indigo-400 dark:to-purple-400">
                Chốt giá hời.
              </span>
            </h1>

            <p className="text-xl md:text-2xl font-bold mb-4 text-gray-800 dark:text-gray-200">
              Sàn đấu giá thời gian thực đỉnh cao.
            </p>

            <p className="max-w-2xl mx-auto lg:mx-0 text-lg text-gray-600 dark:text-gray-400 mb-10 font-medium leading-relaxed">
              Nơi hội tụ những món đồ sưu tầm độc bản từ khắp nơi trên thế giới. Minh bạch - Tốc độ - Uy tín.
            </p>

            <div className="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
              <button
                onClick={() => {
                  document.getElementById('featured-auctions')?.scrollIntoView({ behavior: 'smooth' });
                }}
                className="px-8 py-4 text-lg font-bold bg-blue-600 hover:bg-blue-700 dark:bg-indigo-600 dark:hover:bg-indigo-700 text-white rounded-full shadow-[0_8px_30px_rgb(37,99,235,0.4)] dark:shadow-[0_8px_30px_rgb(79,70,229,0.4)] transition-all transform hover:-translate-y-1 active:scale-95"
              >
                Bắt đầu đấu giá ngay
              </button>
            </div>
          </div>

          <div className="flex-1 w-full max-w-lg lg:max-w-none relative">
            <div className="absolute inset-0 bg-blue-500/10 dark:bg-indigo-500/20 blur-[80px] rounded-full"></div>
            <div className="relative aspect-square w-full">
              <div className="absolute inset-4 rounded-3xl bg-gradient-to-br from-gray-100 to-gray-200 dark:from-white/5 dark:to-white/10 p-8 shadow-2xl border border-white/50 dark:border-white/10 flex items-center justify-center overflow-hidden transform rotate-2 hover:rotate-0 transition-transform duration-500">
                <Image
                  src="/limited_sneaker.png"
                  alt="Limited Sneaker"
                  fill
                  className="object-contain p-8 drop-shadow-[0_20px_50px_rgba(0,0,0,0.3)] dark:drop-shadow-[0_30px_60px_rgba(0,0,0,0.8)] z-10 hover:scale-110 transition-transform duration-700"
                />
              </div>
            </div>
          </div>
        </div>
      </main>

      <section id="featured-auctions" className="py-24 relative bg-gray-50 dark:bg-[#0b0c0e]">
        <div className="container px-6 mx-auto">
          <div className="flex items-end justify-between mb-12">
            <div>
              <h2 className="text-3xl lg:text-4xl font-extrabold mb-3 text-gray-900 dark:text-white">
                Đang Đấu Giá Nóng
              </h2>
              <p className="text-gray-600 dark:text-slate-400 font-medium">Khám phá các bảo vật được săn đón nhất ngay lúc này</p>
            </div>
            <button className="hidden sm:flex items-center gap-2 text-sm font-bold text-blue-600 dark:text-indigo-400 hover:text-blue-800 dark:hover:text-indigo-300 transition-colors group">
              Xem tất cả
              <svg className="w-4 h-4 transform transition-transform group-hover:translate-x-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5">
                <path strokeLinecap="round" strokeLinejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
              </svg>
            </button>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {featuredItems.map((item) => (
              <div
                key={item.id}
                className="group flex flex-col bg-white dark:bg-[#14171d] border border-gray-200 dark:border-white/10 rounded-3xl overflow-hidden hover:border-blue-500/50 dark:hover:border-indigo-500/50 transition-all duration-300 shadow-[0_4px_20px_rgba(0,0,0,0.05)] hover:shadow-[0_20px_40px_rgba(37,99,235,0.15)] hover:-translate-y-2 relative"
              >
                <div className="relative h-72 w-full overflow-hidden bg-gray-100 dark:bg-black/40 flex-shrink-0">
                  <div className="absolute top-4 left-4 z-20 flex flex-col gap-2">
                    <span className="px-3 py-1 bg-white/60 dark:bg-black/40 backdrop-blur-md rounded-full text-xs font-bold text-gray-800 dark:text-white shadow-sm border border-white/20">
                      Hàng Tuyển
                    </span>
                  </div>

                  <div className={`absolute top-4 right-4 z-20 flex items-center gap-1.5 px-3 py-1.5 rounded-full backdrop-blur-md border shadow-sm text-xs font-bold uppercase tracking-wider ${item.isLive ? 'bg-red-500/20 border-red-500/30 text-red-600 dark:text-red-400' : 'bg-orange-500/20 border-orange-500/30 text-orange-600 dark:text-orange-400'}`}>
                    {item.isLive && (
                      <span className="relative flex h-2 w-2 mr-1">
                        <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75"></span>
                        <span className="relative inline-flex rounded-full h-2 w-2 bg-red-500"></span>
                      </span>
                    )}
                    {item.isLive ? "LIVE" : item.timeRemaining}
                  </div>

                  <div className="absolute inset-0 p-6 flex items-center justify-center">
                    <Image
                      src={item.image}
                      alt={item.title}
                      fill
                      className="object-contain p-8 transition-transform duration-500 group-hover:scale-110 drop-shadow-xl"
                    />
                  </div>
                  <div className="absolute inset-0 bg-gradient-to-t from-gray-100/50 dark:from-[#14171d]/80 via-transparent to-transparent opacity-100" />
                </div>

                <div className="p-6 pt-4 flex flex-col flex-grow relative z-10 bg-white dark:bg-[#14171d]">
                  <h3 className="text-xl font-bold mb-3 text-gray-900 dark:text-white group-hover:text-blue-600 dark:group-hover:text-indigo-400 transition-colors line-clamp-2 leading-snug">
                    {item.title}
                  </h3>

                  <div className="mt-auto">
                    <div className="flex items-center justify-between pb-4 border-b border-gray-100 dark:border-white/5 mb-4">
                      <div>
                        <p className="text-xs text-gray-500 dark:text-slate-400 mb-1 font-medium tracking-wide">Giá hiện tại</p>
                        <p className="text-2xl font-extrabold text-green-600 dark:text-green-400 tracking-tight">{item.currentBid}</p>
                      </div>
                      <div className="text-right">
                        <p className="text-xs text-gray-500 dark:text-slate-400 mb-1 font-medium tracking-wide">Tình trạng</p>
                        <p className="text-sm font-bold text-gray-800 dark:text-indigo-300">{item.bids} lượt trả giá</p>
                      </div>
                    </div>

                    <button
                      onClick={() => setIsLoginModalOpen(true)}
                      className="w-full py-4 bg-gray-900 hover:bg-black dark:bg-white dark:hover:bg-gray-200 text-white dark:text-gray-900 text-sm font-extrabold uppercase tracking-widest rounded-xl transition-all shadow-md active:scale-95"
                    >
                      Trả giá ngay
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-10 flex justify-center sm:hidden">
            <button className="w-full py-4 text-center bg-white dark:bg-white/5 border border-gray-200 dark:border-white/10 rounded-xl text-sm font-bold text-gray-900 dark:text-white hover:bg-gray-50 dark:hover:bg-white/10 transition-colors shadow-sm">
              Xem tất cả
            </button>
          </div>
        </div>
      </section>

      <Footer />
      <LoginModal
        isOpen={isLoginModalOpen}
        onClose={() => setIsLoginModalOpen(false)}
        onRegisterClick={() => {
          setIsLoginModalOpen(false);
          setIsRegisterModalOpen(true);
        }}
      />
      <RegisterModal
        isOpen={isRegisterModalOpen}
        onClose={() => setIsRegisterModalOpen(false)}
        onLoginClick={() => {
          setIsRegisterModalOpen(false);
          setIsLoginModalOpen(true);
        }}
      />
    </div>
  );
}
