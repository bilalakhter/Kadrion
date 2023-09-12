import "./globals.css";
import type { Metadata } from "next";
import { Roboto } from "next/font/google";
import Header from "./component/Header";
import Footer from "./component/Footer";
const roboto = Roboto({
  subsets: ["latin"],
  weight: ["400", "500", "700"],
});

export const metadata: Metadata = {
  title: "Kadrion",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={roboto.className}>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  );
}
