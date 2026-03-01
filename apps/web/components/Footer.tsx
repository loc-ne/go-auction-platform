import Link from 'next/link';

export default function Footer() {
    return (
        <footer className="bg-gray-50 dark:bg-[#0b0c0e] py-16 border-t border-gray-200 dark:border-white/10">
            <div className="container mx-auto px-6">
                <div className="grid grid-cols-1 md:grid-cols-4 gap-12 mb-12">
                    <div className="col-span-1 md:col-span-1">
                        <h2 className="text-2xl font-extrabold mb-4 text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 dark:from-indigo-400 dark:to-purple-400">EliteBid</h2>
                        <p className="text-gray-600 dark:text-gray-400 text-sm mb-6 leading-relaxed">
                            Sàn đấu giá thời gian thực nơi bạn tìm thấy những món đồ sưu tầm độc bản, minh bạch, tốc độ và uy tín nhất.
                        </p>
                        <div className="flex gap-4 text-gray-400">
                            <Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400 transition-colors">
                                <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                                    <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z" />
                                </svg>
                            </Link>
                            <Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400 transition-colors">
                                <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z" /></svg>
                            </Link>
                        </div>
                    </div>

                    <div>
                        <h3 className="text-gray-900 dark:text-white font-bold mb-4">Về EliteBid</h3>
                        <ul className="space-y-3 text-sm text-gray-600 dark:text-gray-400">
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Giới thiệu về chúng tôi</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Cách hoạt động</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Bảo hiểm phiên đấu giá</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Blog</Link></li>
                        </ul>
                    </div>

                    <div>
                        <h3 className="text-gray-900 dark:text-white font-bold mb-4">Điều khoản & Hỗ trợ</h3>
                        <ul className="space-y-3 text-sm text-gray-600 dark:text-gray-400">
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Điều khoản đấu giá</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Chính sách bảo mật</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Quy định đổi trả hàng</Link></li>
                            <li><Link href="#" className="hover:text-blue-600 dark:hover:text-indigo-400">Trung tâm trợ giúp</Link></li>
                        </ul>
                    </div>

                    <div>
                        <h3 className="text-gray-900 dark:text-white font-bold mb-4">Tải App EliteBid</h3>
                        <p className="text-sm text-gray-600 dark:text-gray-400 mb-4">Bạn có đồ muốn bán? Tải App ngay để lên sàn trong 30 giây.</p>
                        <div className="flex flex-col gap-3">
                            <button className="flex items-center gap-3 bg-gray-900 dark:bg-white text-white dark:text-gray-900 px-4 py-2.5 rounded-xl hover:bg-gray-800 dark:hover:bg-gray-200 transition-colors">
                                <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 384 512"><path d="M318.7 268.7c-.2-36.7 16.4-64.4 50-84.8-18.8-26.9-47.2-41.7-84.7-44.6-35.5-2.8-74.3 20.7-88.5 20.7-15 0-49.4-19.7-76.4-19.7C63.3 141.2 4 184.8 4 273.5q0 39.3 14.4 81.2c12.8 36.7 59 126.7 107.2 125.2 25.2-.6 43-17.9 75.8-17.9 31.8 0 48.3 17.9 76.4 17.9 48.6-.7 90.4-82.5 102.6-119.3-65.2-30.7-61.7-90-61.7-91.9zm-56.6-164.2c27.3-32.4 24.8-61.9 24-72.5-24.1 1.4-52 16.4-67.9 34.9-17.5 19.8-27.8 44.3-25.6 71.9 26.1 2 49.9-11.4 69.5-34.3z" /></svg>
                                <div className="text-left">
                                    <div className="text-[10px] leading-none mb-1">Download on the</div>
                                    <div className="text-sm font-bold leading-none">App Store</div>
                                </div>
                            </button>
                            <button className="flex items-center gap-3 bg-gray-900 dark:bg-white text-white dark:text-gray-900 px-4 py-2.5 rounded-xl hover:bg-gray-800 dark:hover:bg-gray-200 transition-colors">
                                <svg className="w-6 h-6" viewBox="0 0 512 512" fill="currentColor"><path d="M325.3 234.3L104.6 13l280.8 161.2-60.1 60.1zM47 0C34 6.8 25.3 19.2 25.3 35.3v441.3c0 16.1 8.7 28.5 21.7 35.3l256.6-256L47 0zm425.2 225.6l-58.9-34.1-65.7 64.5 65.7 64.5 60.1-34.1c18-14.3 18-46.5-1.2-60.8zM104.6 499l280.8-161.2-60.1-60.1L104.6 499z" /></svg>
                                <div className="text-left">
                                    <div className="text-[10px] leading-none mb-1">GET IT ON</div>
                                    <div className="text-sm font-bold leading-none">Google Play</div>
                                </div>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </footer>
    );
}
