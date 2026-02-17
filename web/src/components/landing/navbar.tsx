"use client"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Github, ArrowRight } from "lucide-react"
import { motion } from "framer-motion"

export function Navbar() {
    return (
        <motion.nav
            initial={{ y: -20, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ duration: 0.5 }}
            className="fixed top-0 w-full z-50 border-b border-white/[0.06] bg-black/60 backdrop-blur-2xl backdrop-saturate-150"
        >
            <div className="max-w-6xl mx-auto px-6 h-14 flex items-center justify-between">
                {/* Logo */}
                <Link href="/" className="flex items-center gap-2.5 group">
                    <div className="h-7 w-7 rounded-md bg-white flex items-center justify-center group-hover:bg-white/90 transition-colors">
                        <span className="text-black font-bold text-xs">O</span>
                    </div>
                    <span className="font-semibold text-[15px] tracking-tight text-white/90">Orion</span>
                </Link>

                {/* Center Nav Links */}
                <div className="hidden md:flex items-center gap-1">
                    <Link
                        href="#features"
                        className="px-3 py-1.5 text-[13px] text-white/40 hover:text-white/80 transition-colors rounded-md hover:bg-white/[0.04]"
                    >
                        Features
                    </Link>
                    <Link
                        href="#how-it-works"
                        className="px-3 py-1.5 text-[13px] text-white/40 hover:text-white/80 transition-colors rounded-md hover:bg-white/[0.04]"
                    >
                        How It Works
                    </Link>
                    <Link
                        href="#install"
                        className="px-3 py-1.5 text-[13px] text-white/40 hover:text-white/80 transition-colors rounded-md hover:bg-white/[0.04]"
                    >
                        Install
                    </Link>
                </div>

                {/* Right Actions */}
                <div className="flex items-center gap-3">
                    <Link href="https://github.com/TanmayDabhade/orion" target="_blank">
                        <Button variant="ghost" size="icon" className="h-8 w-8 text-white/30 hover:text-white/80 hover:bg-white/[0.06]">
                            <Github className="h-4 w-4" />
                        </Button>
                    </Link>
                    <Link href="#install">
                        <Button className="h-8 px-3.5 text-[13px] bg-white text-black hover:bg-white/90 font-medium rounded-lg relative overflow-hidden shine-hover">
                            Install v0.7.0
                            <ArrowRight className="ml-1.5 h-3.5 w-3.5" />
                        </Button>
                    </Link>
                </div>
            </div>
        </motion.nav>
    )
}
