import { ArrowRightLeft } from 'lucide-react'
import { useNavigate } from 'react-router'
import { appname } from '../global'

export default function Header() {
  const navigate = useNavigate()

  function handleGoHome() {
    navigate('/')
  }

  return (
    <>
      <header className='bg-white shadow-sm'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4'>
          <div className='flex items-center justify-between'>
            <div className='flex items-center'>
              <ArrowRightLeft className='h-8 w-8 text-indigo-600' />
              <h1
                onClick={handleGoHome}
                className='ml-2 text-2xl font-bold text-gray-900 cursor-pointer'
              >
                {appname}
              </h1>
            </div>
          </div>
        </div>
      </header>
    </>
  )
}
