import Prelude

mass :: Int -> Int
mass x = x `div` 3 - 2

main = interact $ show . sum . map mass . map read . words
