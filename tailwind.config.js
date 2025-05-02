/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./template.html", // Directly specify the file
        // Or use a pattern if you have multiple HTML files
        // "./*.html",
        // Add paths to other files that might contain Tailwind classes
        // "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {},
    },
    plugins: [],
}