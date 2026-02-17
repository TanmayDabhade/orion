import Link from "next/link"
import { Github } from "lucide-react"

export function Footer() {
    return (
        <footer className="border-t border-white/[0.04] py-10">
            <div className="max-w-6xl mx-auto px-6">
                <div className="flex flex-col md:flex-row items-center justify-between gap-4">
                    {/* Logo + Copyright */}
                    <div className="flex items-center gap-3">
                        <div className="h-6 w-6 rounded-md bg-white flex items-center justify-center">
                            <span className="text-black font-bold text-[10px]">O</span>
                        </div>
                        <span className="text-[13px] text-white/25">
                            &copy; {new Date().getFullYear()} Orion. Built by{" "}
                            <Link href="https://github.com/TanmayDabhade" target="_blank" className="text-white/40 hover:text-white/70 transition-colors">
                                Tanmay Dabhade
                            </Link>
                        </span>
                    </div>

                    {/* Links */}
                    <div className="flex items-center gap-5 text-[13px] text-white/25">
                        <Link href="https://github.com/TanmayDabhade/orion" target="_blank" className="hover:text-white/60 transition-colors flex items-center gap-1.5">
                            <Github className="h-3.5 w-3.5" />
                            GitHub
                        </Link>
                        <Link href="https://github.com/TanmayDabhade/orion/releases" target="_blank" className="hover:text-white/60 transition-colors">
                            Releases
                        </Link>
                        <Link href="https://github.com/TanmayDabhade/orion#contributing" target="_blank" className="hover:text-white/60 transition-colors">
                            Contribute
                        </Link>
                    </div>
                </div>
            </div>
        </footer>
    )
}
