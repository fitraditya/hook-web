import { ArrowRight } from 'lucide-react'
import { useEffect } from 'react'
import { useNavigate } from 'react-router'
import { v4 as uuid } from 'uuid'
import Footer from './components/Footer'
import Header from './components/Header'
import { appname } from './global'

export default function Home() {
  const navigate = useNavigate()

  useEffect(() => {
    document.title = appname
  })

  function createBin() {
    const slug = uuid()
    const url = `/${slug}/inspect`
    navigate(url)
  }

  return (
    <>
      <Header></Header>
      <section className='pt-20 pb-32 px-4 text-center'>
        <div className='max-w-3xl mx-auto'>
          <h1 className='text-4xl sm:text-5xl font-bold text-gray-900 mb-6'>
            Debug and Inspect HTTP Requests with Ease
          </h1>
          <p className='text-xl text-gray-600 mb-8'>
            HookWeb gives you a URL that collects requests you send it and
            lets you inspect them in a human-friendly way. Perfect for debugging
            webhooks and HTTP clients.
          </p>
          <div className='flex flex-col sm:flex-row items-center justify-center gap-4'>
            <button
              onClick={createBin}
              className='w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 cursor-pointer'
            >
              Get Started
              <ArrowRight className='ml-2 h-5 w-5' />
            </button>
          </div>
        </div>
      </section>
      <Footer></Footer>
    </>
  )
}
