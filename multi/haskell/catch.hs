import Data.Char (toLower)
import Data.List (isInfixOf)

spamWords :: [String]
spamWords =
	[ "free", "winner", "win", "prize", "cash", "money", "offer"
	, "buy now", "click here", "urgent", "congratulations", "viagra"
	]

isSpam :: String -> Bool
isSpam message =
	length (filter (`isInfixOf` text) spamWords) >= 2
	where
		text = map toLower message

main :: IO ()
main = do
	message <- getContents
	putStrLn $ if isSpam message then "SPAM" else "HAM"
