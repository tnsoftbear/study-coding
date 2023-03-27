import Head from "next/head";
import Header from "./Header";
import Footer from "./Footer";
import { ReactNode } from "react";
import styles from '../styles/Layout.module.css'

interface LayoutProps {
  title: string;
  keywords: string;
  description: string;
  children: ReactNode;
}

export default function Layout({
  title,
  keywords,
  description,
  children,
}: LayoutProps) {
  return (
    <div>
      <Head>
        <title>{title}</title>
        <meta name="description" content={description} />
        <meta name="keywords" content={keywords} />
      </Head>

      <Header />
      <div className={styles.container}>{children}</div>
      <Footer />
    </div>
  );
}

Layout.defaultProps = {
  title: "DJ Events | Find parties",
  description: "Find lastest events",
  keywords: "music, dj, mc",
};
