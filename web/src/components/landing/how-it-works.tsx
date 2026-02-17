"use client"

import { motion } from "framer-motion"

const steps = [
    {
        step: "01",
        title: "Add shortcuts",
        description: "Map any URL or command to a short alias.",
        code: [
            { prompt: true, text: 'o add d2l "open https://d2l.msu.edu"' },
            { prompt: false, text: "✓ Shortcut 'd2l' added" },
        ],
    },
    {
        step: "02",
        title: "Launch anything",
        description: "Apps, URLs, and shortcuts. One command.",
        code: [
            { prompt: true, text: "o chrome" },
            { prompt: false, text: "→ Opening Google Chrome.app" },
            { prompt: true, text: "o d2l" },
            { prompt: false, text: "→ Opening https://d2l.msu.edu" },
        ],
    },
    {
        step: "03",
        title: "Stay in control",
        description: "See everything at a glance. Manage your workflow.",
        code: [
            { prompt: true, text: "o list" },
            { prompt: false, text: "d2l     → open https://d2l.msu.edu" },
            { prompt: false, text: "mail    → open https://mail.google.com" },
            { prompt: false, text: "github  → open https://github.com" },
        ],
    },
]

export function HowItWorks() {
    return (
        <section id="how-it-works" className="relative py-28 overflow-hidden border-t border-white/[0.04]">
            <div className="max-w-6xl mx-auto px-6">
                {/* Section Header */}
                <motion.div
                    initial={{ opacity: 0, y: 20 }}
                    whileInView={{ opacity: 1, y: 0 }}
                    viewport={{ once: true }}
                    className="text-center mb-16"
                >
                    <p className="text-[13px] text-white/30 uppercase tracking-widest mb-3 font-medium">How It Works</p>
                    <h2 className="text-3xl md:text-4xl font-bold tracking-tight text-white mb-4">
                        Three commands. That&apos;s it.
                    </h2>
                    <p className="text-white/35 text-base max-w-lg mx-auto">
                        No config files. No setup wizards. Just install and go.
                    </p>
                </motion.div>

                {/* Steps */}
                <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    {steps.map((step, index) => (
                        <motion.div
                            key={index}
                            initial={{ opacity: 0, y: 20 }}
                            whileInView={{ opacity: 1, y: 0 }}
                            viewport={{ once: true }}
                            transition={{ delay: index * 0.1 }}
                            className="relative"
                        >
                            <div className="rounded-2xl border border-white/[0.06] bg-white/[0.02] overflow-hidden">
                                {/* Step Header */}
                                <div className="p-6 pb-4">
                                    <span className="text-[12px] font-mono text-white/20 mb-2 block">{step.step}</span>
                                    <h3 className="text-lg font-semibold text-white mb-1">{step.title}</h3>
                                    <p className="text-[13px] text-white/35">{step.description}</p>
                                </div>

                                {/* Code Block */}
                                <div className="mx-4 mb-4 rounded-xl bg-[#050505] border border-white/[0.06] p-4 font-mono text-[12px] leading-loose">
                                    {step.code.map((line, i) => (
                                        <div key={i} className={line.prompt ? "text-white/60" : "text-white/25"}>
                                            {line.prompt && <span className="text-white/20 mr-1">$</span>}
                                            {!line.prompt && <span className="ml-2.5">{line.text}</span>}
                                            {line.prompt && <span>{line.text}</span>}
                                        </div>
                                    ))}
                                </div>
                            </div>
                        </motion.div>
                    ))}
                </div>
            </div>
        </section>
    )
}
