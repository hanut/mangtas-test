package main

// Type definition for the WordCountResult array
type WCRList []*WordCountResult

// implementation of Len() for sort
func (wcrl WCRList) Len() int { return len(wcrl) }

// implementation of Swap() for sort
func (wcrl WCRList) Swap(a, b int) { wcrl[a], wcrl[b] = wcrl[b], wcrl[a] }

type ByCount struct{ WCRList }

// implementation of Swap() for sort
func (wcrl ByCount) Less(a, b int) bool { return wcrl.WCRList[a].Count < wcrl.WCRList[b].Count }
