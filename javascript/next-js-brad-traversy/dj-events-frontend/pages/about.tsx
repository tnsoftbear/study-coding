import Link from 'next/link'
import Layout from '@/components/Layout'

export default function AboutPage() {
  return (
    <Layout title="About DJ">
      <h1>About</h1>
      <p>This is example project based on the next.js course</p>
      <p>Thanks to teacher</p>
      <Link href='/'>Home</Link>
    </Layout>
  )
}
