package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/taan02991/blog/x/blog/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateComment(ctx sdk.Context, comment types.Comment) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.CommentPrefix + comment.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	store.Set(key, value)
}

func listComment(ctx sdk.Context, k Keeper) ([]byte, error) {
  var commentList []types.Comment
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.CommentPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var comment types.Comment
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &comment)
    commentList = append(commentList, comment)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, commentList)
  return res, nil
}