"use client"

import { Button } from "@/components/ui/button"
import { ArrowRight, Copy, Check } from "lucide-react"
import { motion } from "framer-motion"
import { useState } from "react"

export function Install() {
    const [copied, setCopied] = useState(false)
    const cmd = "curl -fsSL https://github.com/TanmayDabhade/orion/releases/latest/download/install.sh | sh"

    const copy = () => {
        navigator.clipboard.writeText(cmd)
        setCopied(true)
        setTimeout(() => setCopied(false), 2000)
    }

    return (
        <section id="install" className="relative py-28 overflow-hidden border-t border-white/[0.04]">
            <div className="max-w-3xl mx-auto px-6 text-center">
                <motion.div
                    initial={{ opacity: 0, y: 20 }}
                    whileInView={{ opacity: 1, y: 0 }}
                    viewport={{ once: true }}
                >
                    <h2 className="text-3xl md:text-4xl font-bold tracking-tight text-white mb-4">
                        Ready to try Orion?
                    </h2>
                    <p className="text-white/35 text-base mb-10 max-w-md mx-auto">
                        One command to install. Works with Bash, Zsh, and Fish.
                    </p>

                    {/* Install Command */}
                    <div
                        onClick={copy}
                        className="group relative mx-auto max-w-2xl rounded-2xl border border-white/[0.08] bg-[#050505] p-5 font-mono text-[13px] text-left cursor-pointer hover:border-white/[0.15] transition-colors"
                    >
                        <div className="flex items-start justify-between gap-4">
                            <div className="text-white/50 break-all leading-relaxed">
                                <span className="text-white/20">$ </span>
                                {cmd}
                            </div>
                            <button className="shrink-0 mt-0.5 p-1.5 rounded-md hover:bg-white/[0.06] transition-colors">
                                {copied ? (
                                    <Check className="h-4 w-4 text-white" />
                                ) : (
                                    <Copy className="h-4 w-4 text-white/25 group-hover:text-white/50 transition-colors" />
                                )}
                            </button>
                        </div>
                    </div>

                    <div className="mt-8 flex flex-col sm:flex-row items-center justify-center gap-3">
                        <Button
                            variant="outline"
                            size="lg"
                            className="h-11 px-6 bg-transparent border-white/[0.08] hover:bg-white/[0.04] hover:border-white/[0.12] text-white/50 text-sm rounded-xl"
                            asChild
                        >
                            <a href="https://github.com/TanmayDabhade/orion" target="_blank">
                                View on GitHub
                            </a>
                        </Button>
                    </div>
                </motion.div>
            </div>
        </section>
    )
}
