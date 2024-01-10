import Data.Char
import Data.Map
import qualified Data.Map as Map
-- How many steps are required?

data Island = Island { name :: String, leff :: String, righ :: String } deriving Show

-- trie like data structure
data Tree = Node (Map Char Tree) | Leaf Island | EmptyTree
-- find the node pertaining to the string if it exists (go backwards)
find :: Tree -> String -> Island

-- given an empty tree
find EmptyTree _ = error "searching an empty tree"

-- a leaf and an empty string
find (Leaf n) "" = n
-- a leaf and a nonempty string
find (Leaf n) x = error "asking for an island that doesn't exist" ++ x

-- single character string
find (Node m) [x] = case findWithDefault (error "asking for an island that doesn't exist") x m of
        Node n -> error "should be at leaf but at a node"
        Leaf n -> n
        other -> other

-- >= 2 character string
find (Node m) s = let newIsland = findWithDefault (error "asking for an island that doesn't exist") (last s) m in
        findWithDefault (error "asking for an island that doesn't exist") (last $ init s) newIsland

makeTree :: [Island] -> Tree -> Tree
makeTree [] t = t
makeTree (i:is) t = let newTree = insertIntoTree i t in
        makeTree is newTree

insertIntoTree :: Island -> Tree -> Tree 
insertIntoTree i EmptyTree = let namae = name n in
        let lastTree = Leaf i
            thirdTree  = Island (singleton (Prelude.take 1 namae) lastTree)
            secondTree = Island (singleton (last $ Prelude.take 2 namae) thirdTree)
            in Island (singleton (last namae) secondTree)
        

-- given a character and a string, split the string up based on this character
split :: Char -> String -> [String]
split d s = [takeWhile (\c -> c /= d) s, Prelude.drop 1 $ Prelude.dropWhile (\c -> c /= d) s]

stripWhiteSpace :: String -> String
stripWhiteSpace s = (Prelude.dropWhile isSpace . reverse . Prelude.dropWhile isSpace . reverse) s

-- precondition is that the string represents a tuple
removeTupleBrackets :: String -> String
removeTupleBrackets s = Prelude.drop 1 $ Prelude.take (length s - 1) s

-- given a string that is a line representing a node
line2Island :: String -> Island
line2Island rawString =
  let 
    rawList = Prelude.map stripWhiteSpace $ Main.split '=' rawString
    rawOutgoingIslands = rawList !! 1
    outgoingIslands = Prelude.map stripWhiteSpace $ Main.split ',' $ removeTupleBrackets $ stripWhiteSpace rawOutgoingIslands
        in Island (rawList !! 0) (outgoingIslands !! 0) (outgoingIslands !! 1)

printList :: Show a => [a] -> IO ()
printList [] = return ()
printList (x:xs) = do
        print x
        printList xs

-- given the current number of steps, the current islands, the instructions from
-- which we're up to and the tree, determine the number of steps it took to
-- get to the terminating nodes
followInstructions :: Int -> [Island] -> String -> Tree -> IO Int
followInstructions acc currs "" t = return acc
followInstructions acc currs (x:xs) t = do
        -- get the string of the new islands we'll be at
        let newCurrsString = Prelude.map (\n -> if (last $ name n) == 'L' then leff n else righ n) currs
        -- find the nodes corresponding to the strings
        let newCurrs = Prelude.map (\s -> find t s) newCurrsString
        let isDone = all (
                \n -> case n of
                        Nothing -> False
                        Just x -> if (last $ name x) == 'Z' then True else False) newCurrs

        print acc
        printList newCurrs
        
        if isDone then do return (acc + 1) else
                followInstructions (acc + 1) newCurrs xs t
                        
-- find the node to go to then recurse on this node
main :: IO ()
main = do
        inp <- lines <$> readFile "input.txt"

-- get the first line; instructions
        let inst = head inp
        
-- get the list of nodes
        let islands = Prelude.map (line2Island) $ Prelude.drop 2 inp
-- construct a node tree using the list of nodesnodes
        let islandsTree = makeTree islands EmptyTree
         
-- using the tree get the nodetree (or list) that pertains to the list of
-- starts
        let starts = Prelude.filter (\n -> (last $ name n) == 'A') islands

-- for each instruction we will 
        ans <- followInstructions 0 starts (cycle inst) islandsTree
        print ans
