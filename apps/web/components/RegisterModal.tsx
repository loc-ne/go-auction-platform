import { useState } from 'react';
import { useAuth } from '../context/AuthContext';

export default function RegisterModal({ isOpen, onClose, onLoginClick }: { isOpen: boolean; onClose: () => void; onLoginClick?: () => void }) {
    const [fullName, setFullName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    if (!isOpen) return null;

    const API_URL = process.env.NEXT_PUBLIC_API_URL;

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setIsLoading(true);

        try {
            const res = await fetch(`${API_URL}/auth/register`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password, full_name: fullName }),
            });

            const data = await res.json();

            if (res.ok && data.success) {
                if (onLoginClick) {
                    onClose();
                    onLoginClick();
                }
            } else {
                setError(data.message || 'Đăng ký thất bại.');
            }
        } catch (err) {
            setError('Đã có lỗi xảy ra. Mời thử lại sau.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="fixed inset-0 z-[100] flex items-center justify-center">
            <div
                className="absolute inset-0 bg-gray-900/40 dark:bg-black/60 backdrop-blur-md transition-opacity"
                onClick={onClose}
            ></div>

            <div className="relative w-full max-w-md bg-white dark:bg-[#14171d] rounded-3xl shadow-2xl overflow-hidden border border-gray-100 dark:border-white/10 p-8 transform transition-all animate-in fade-in zoom-in duration-300">
                <div className="absolute top-4 right-4">
                    <button onClick={onClose} className="p-2 text-gray-400 hover:text-gray-600 dark:hover:text-white rounded-full bg-gray-50 dark:bg-white/5 hover:bg-gray-100 dark:hover:bg-white/10 transition-colors">
                        <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>

                <h2 className="text-3xl font-extrabold text-center mb-8 text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 dark:from-indigo-400 dark:to-purple-400">
                    Đăng ký tài khoản
                </h2>

                <form className="space-y-5" onSubmit={handleRegister}>
                    {error && (
                        <div className="p-3 mb-4 text-sm text-red-500 bg-red-100/50 dark:bg-red-500/10 border border-red-500/20 rounded-xl">
                            {error}
                        </div>
                    )}
                    <div className="space-y-1.5">
                        <div className="relative">
                            <div className="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none">
                                <svg className="w-5 h-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                                    <path fillRule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clipRule="evenodd" />
                                </svg>
                            </div>
                            <input
                                type="text"
                                placeholder="Họ và tên"
                                className="w-full pl-11 pr-4 py-3 bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-indigo-500 transition-all font-medium placeholder-gray-400"
                                required
                                value={fullName}
                                onChange={(e) => setFullName(e.target.value)}
                            />
                        </div>
                    </div>

                    <div className="space-y-1.5">
                        <div className="relative">
                            <div className="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none">
                                <svg className="w-5 h-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                                    <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                                    <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                                </svg>
                            </div>
                            <input
                                type="email"
                                placeholder="Email của bạn"
                                className="w-full pl-11 pr-4 py-3 bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-indigo-500 transition-all font-medium placeholder-gray-400"
                                required
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
                        </div>
                    </div>

                    <div className="space-y-1.5">
                        <div className="relative">
                            <div className="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none">
                                <svg className="w-5 h-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                                    <path fillRule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clipRule="evenodd" />
                                </svg>
                            </div>
                            <input
                                type="password"
                                placeholder="Mật khẩu"
                                className="w-full pl-11 pr-4 py-3 bg-gray-50 dark:bg-black/20 border border-gray-200 dark:border-white/10 rounded-xl text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-indigo-500 transition-all font-medium placeholder-gray-400"
                                required
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>
                    </div>

                    <button
                        type="submit"
                        disabled={isLoading}
                        className={`w-full py-4 text-white font-extrabold text-base rounded-xl transition-all shadow-lg active:scale-95 mt-6 flex justify-center items-center ${isLoading
                            ? 'bg-blue-400 dark:bg-indigo-400 cursor-not-allowed shadow-none'
                            : 'bg-blue-600 hover:bg-blue-700 dark:bg-indigo-500 dark:hover:bg-indigo-600 shadow-blue-500/30 dark:shadow-indigo-500/25'
                            }`}
                    >
                        {isLoading ? (
                            <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                        ) : (
                            'Tạo tài khoản'
                        )}
                    </button>
                </form>

                <p className="mt-8 text-center text-sm font-medium text-gray-600 dark:text-gray-400">
                    Đã có tài khoản?{' '}
                    <button
                        onClick={() => {
                            onClose();
                            if (onLoginClick) onLoginClick();
                        }}
                        className="text-blue-600 dark:text-indigo-400 font-bold hover:underline bg-transparent border-none p-0 cursor-pointer"
                    >
                        Đăng nhập ngay
                    </button>
                </p>
            </div>
        </div>
    );
}
