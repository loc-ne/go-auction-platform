import Link from 'next/link';
import { useAuth } from '../context/AuthContext';

export default function Navbar({ onLoginClick, onRegisterClick }: { onLoginClick: () => void; onRegisterClick?: () => void }) {
    const { user, isLoading, logout } = useAuth();

    return (
        <nav className="fixed top-0 left-0 right-0 z-50 flex items-center justify-between px-6 py-4 bg-white/80 dark:bg-[#0f1115]/80 backdrop-blur-xl border-b border-gray-200 dark:border-white/5">
            <div className="flex items-center gap-8">
                <Link href="/" className="text-2xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 dark:from-indigo-400 dark:to-purple-400">
                    EliteBid
                </Link>
                <div className="hidden md:flex items-center gap-6 text-sm font-semibold text-gray-700 dark:text-gray-300">
                    <Link href="#" className="hover:text-blue-600 dark:hover:text-white transition-colors">Khám phá</Link>
                    <Link href="#" className="hover:text-blue-600 dark:hover:text-white transition-colors">Đang hot</Link>
                    <Link href="#" className="hover:text-blue-600 dark:hover:text-white transition-colors">Về chúng tôi</Link>
                </div>
            </div>

            <div className="hidden lg:flex flex-1 max-w-md mx-8 items-center bg-gray-100 dark:bg-white/5 border border-transparent dark:border-white/10 rounded-full px-4 py-2 hover:bg-gray-200 dark:hover:bg-white/10 transition-colors focus-within:border-blue-500 focus-within:ring-1 focus-within:ring-blue-500 focus-within:bg-white dark:focus-within:bg-white/5">
                <svg className="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input
                    type="text"
                    placeholder="Tìm Jordan 1, Rolex..."
                    className="w-full bg-transparent border-none outline-none px-3 text-sm text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400"
                />
            </div>

            <div className="flex items-center gap-3">
                {isLoading ? (
                    <div className="w-8 h-8 border-2 border-blue-600 dark:border-indigo-400 border-t-transparent rounded-full animate-spin"></div>
                ) : !user ? (
                    <>
                        <button onClick={onRegisterClick} className="px-5 py-2 text-sm font-bold text-blue-600 dark:text-indigo-400 border-2 border-blue-600 dark:border-indigo-500 rounded-full hover:bg-blue-50 dark:hover:bg-indigo-500/10 transition-colors hidden sm:block">
                            Đăng ký
                        </button>
                        <button onClick={onLoginClick} className="px-5 py-2.5 text-sm font-bold bg-blue-600 dark:bg-indigo-600 text-white rounded-full hover:bg-blue-700 dark:hover:bg-indigo-700 shadow-[0_4px_14px_0_rgba(37,99,235,0.39)] dark:shadow-[0_4px_14px_0_rgba(79,70,229,0.39)] transition-all hover:shadow-[0_6px_20px_rgba(37,99,235,0.23)] dark:hover:shadow-[0_6px_20px_rgba(79,70,229,0.23)] active:scale-95">
                            Đăng nhập
                        </button>
                    </>
                ) : (
                    <div className="flex items-center gap-4">
                        <button onClick={logout} title="Đăng xuất" className="relative p-2 text-gray-600 dark:text-gray-300 hover:text-red-600 dark:hover:text-red-400 transition-colors">
                            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                            </svg>
                        </button>
                        <button className="relative p-2 text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-white transition-colors">
                            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                            </svg>
                            <span className="absolute top-1 right-1 w-2.5 h-2.5 bg-red-500 border-2 border-white dark:border-[#0f1115] rounded-full"></span>
                        </button>
                        <div className="text-right hidden sm:block">
                            <p className="text-[10px] text-gray-500 dark:text-gray-400 font-medium leading-tight">Xin chào</p>
                            <p className="text-sm font-bold text-gray-900 dark:text-white leading-tight">{user.fullName || user.email}</p>
                        </div>
                        <div className="w-10 h-10 rounded-full bg-gradient-to-tr from-blue-500 to-purple-500 border-2 border-white dark:border-[#0f1115] shadow-sm flex items-center justify-center text-white font-bold cursor-pointer hover:opacity-90 transition-opacity">
                            {user.fullName ? user.fullName.charAt(0).toUpperCase() : user.email.charAt(0).toUpperCase()}
                        </div>
                    </div>
                )}
            </div>
        </nav>
    );
}
