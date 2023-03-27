import Link from 'next/link'
import styles from '../styles/Footer.module.css'

export default function Header() {
  return (
    <footer className={styles.footer}>
        <p>Copyright &copy; tn-softbear</p>
        <p>
            <Link href='/about'>About me</Link>
        </p>
    </footer>
  )
}
