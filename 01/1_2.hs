import Prelude

mass :: Int -> Int
mass x = x `div` 3 - 2

mass' :: Int -> Int
mass' x
  | m <= 0    = 0
  | otherwise = m + mass' m
  where m = mass x

main = interact $ show . sum . map mass' . map read . words
