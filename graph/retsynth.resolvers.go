package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
)

// UniqueMetabolicClusters is the resolver for the uniqueMetabolicClusters field.
func (r *queryResolver) UniqueMetabolicClusters(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: UniqueMetabolicClusters - uniqueMetabolicClusters"))
}

// ModelIDsFromCluster is the resolver for the modelIDsFromCluster field.
func (r *queryResolver) ModelIDsFromCluster(ctx context.Context, cluster *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ModelIDsFromCluster - modelIDsFromCluster"))
}

// AllModelIDs is the resolver for the allModelIDs field.
func (r *queryResolver) AllModelIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: AllModelIDs - allModelIDs"))
}

// OrganismName is the resolver for the organismName field.
func (r *queryResolver) OrganismName(ctx context.Context, organismID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: OrganismName - organismName"))
}

// OrganismID is the resolver for the organismID field.
func (r *queryResolver) OrganismID(ctx context.Context, organismName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: OrganismID - organismID"))
}

// CompoundID is the resolver for the compoundID field.
func (r *queryResolver) CompoundID(ctx context.Context, compoundName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundID - compoundID"))
}

// LikeCompoundID is the resolver for the likeCompoundID field.
func (r *queryResolver) LikeCompoundID(ctx context.Context, compoundName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: LikeCompoundID - likeCompoundID"))
}

// CompoundIDFromInchi is the resolver for the compoundIDFromInchi field.
func (r *queryResolver) CompoundIDFromInchi(ctx context.Context, inchi *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundIDFromInchi - compoundIDFromInchi"))
}

// CompoundInchi is the resolver for the compoundInchi field.
func (r *queryResolver) CompoundInchi(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundInchi - compoundInchi"))
}

// CompoundNameFromInchi is the resolver for the compoundNameFromInchi field.
func (r *queryResolver) CompoundNameFromInchi(ctx context.Context, inchi *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundNameFromInchi - compoundNameFromInchi"))
}

// CompoundName is the resolver for the compoundName field.
func (r *queryResolver) CompoundName(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundName - compoundName"))
}

// CompoundCompartment is the resolver for the compoundCompartment field.
func (r *queryResolver) CompoundCompartment(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundCompartment - compoundCompartment"))
}

// ReactionName is the resolver for the reactionName field.
func (r *queryResolver) ReactionName(ctx context.Context, reactionID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ReactionName - reactionName"))
}

// ReactionID is the resolver for the reactionID field.
func (r *queryResolver) ReactionID(ctx context.Context, reactionName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ReactionID - reactionID"))
}

// ReactionIDsFromCompound is the resolver for the reactionIDsFromCompound field.
func (r *queryResolver) ReactionIDsFromCompound(ctx context.Context, compoundID *string, isProduct *bool) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionIDsFromCompound - reactionIDsFromCompound"))
}

// ReactionSpecies is the resolver for the reactionSpecies field.
func (r *queryResolver) ReactionSpecies(ctx context.Context, reactionID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionSpecies - reactionSpecies"))
}

// ReactantCompoundIDs is the resolver for the reactantCompoundIDs field.
func (r *queryResolver) ReactantCompoundIDs(ctx context.Context, reactionID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactantCompoundIDs - reactantCompoundIDs"))
}

// GetReactionsWithProduct is the resolver for the getReactionsWithProduct field.
func (r *queryResolver) GetReactionsWithProduct(ctx context.Context, compoundID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: GetReactionsWithProduct - getReactionsWithProduct"))
}

// ModelCompounds is the resolver for the modelCompounds field.
func (r *queryResolver) ModelCompounds(ctx context.Context, modelID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ModelCompounds - modelCompounds"))
}

// CompoundIDs is the resolver for the compoundIDs field.
func (r *queryResolver) CompoundIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundIDs - compoundIDs"))
}

// CompoundInchiStrings is the resolver for the compoundInchiStrings field.
func (r *queryResolver) CompoundInchiStrings(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundInchiStrings - compoundInchiStrings"))
}

// ModelReactions is the resolver for the modelReactions field.
func (r *queryResolver) ModelReactions(ctx context.Context, modelID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ModelReactions - modelReactions"))
}

// ReactionIDs is the resolver for the reactionIDs field.
func (r *queryResolver) ReactionIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionIDs - reactionIDs"))
}

// ReactionReversibility is the resolver for the reactionReversibility field.
func (r *queryResolver) ReactionReversibility(ctx context.Context, reactionID *string, modelID *string) (*bool, error) {
	panic(fmt.Errorf("not implemented: ReactionReversibility - reactionReversibility"))
}

// ReactionReversibilityGlobal is the resolver for the reactionReversibilityGlobal field.
func (r *queryResolver) ReactionReversibilityGlobal(ctx context.Context, reactionID *string) (*bool, error) {
	panic(fmt.Errorf("not implemented: ReactionReversibilityGlobal - reactionReversibilityGlobal"))
}

// ReactionGeneAssociations is the resolver for the reactionGeneAssociations field.
func (r *queryResolver) ReactionGeneAssociations(ctx context.Context, reactionID *string, modelID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionGeneAssociations - reactionGeneAssociations"))
}

// ReactionProteinAssociations is the resolver for the reactionProteinAssociations field.
func (r *queryResolver) ReactionProteinAssociations(ctx context.Context, reactionID *string, modelID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionProteinAssociations - reactionProteinAssociations"))
}

// Stoichiometry is the resolver for the stoichiometry field.
func (r *queryResolver) Stoichiometry(ctx context.Context, reactionID *string, compoundID *string, isProduct *bool) (*float64, error) {
	panic(fmt.Errorf("not implemented: Stoichiometry - stoichiometry"))
}

// ReactionCatalysts is the resolver for the reactionCatalysts field.
func (r *queryResolver) ReactionCatalysts(ctx context.Context, reactionID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionCatalysts - reactionCatalysts"))
}

// CompartmentID is the resolver for the compartmentID field.
func (r *queryResolver) CompartmentID(ctx context.Context, compartmentName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompartmentID - compartmentID"))
}

// ReactionSolvents is the resolver for the reactionSolvents field.
func (r *queryResolver) ReactionSolvents(ctx context.Context, reactionID *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionSolvents - reactionSolvents"))
}

// ReactionTemperature is the resolver for the reactionTemperature field.
func (r *queryResolver) ReactionTemperature(ctx context.Context, reactionID *string) (*float64, error) {
	panic(fmt.Errorf("not implemented: ReactionTemperature - reactionTemperature"))
}

// ReactionPressure is the resolver for the reactionPressure field.
func (r *queryResolver) ReactionPressure(ctx context.Context, reactionID *string) (*float64, error) {
	panic(fmt.Errorf("not implemented: ReactionPressure - reactionPressure"))
}

// ReactionTime is the resolver for the reactionTime field.
func (r *queryResolver) ReactionTime(ctx context.Context, reactionID *string) (*float64, error) {
	panic(fmt.Errorf("not implemented: ReactionTime - reactionTime"))
}

// ReactionYield is the resolver for the reactionYield field.
func (r *queryResolver) ReactionYield(ctx context.Context, reactionID *string) (*float64, error) {
	panic(fmt.Errorf("not implemented: ReactionYield - reactionYield"))
}

// ReactionReference is the resolver for the reactionReference field.
func (r *queryResolver) ReactionReference(ctx context.Context, reactionID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ReactionReference - reactionReference"))
}

// ReactionByType is the resolver for the reactionByType field.
func (r *queryResolver) ReactionByType(ctx context.Context, reactionType *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionByType - reactionByType"))
}

// ReactionType is the resolver for the reactionType field.
func (r *queryResolver) ReactionType(ctx context.Context, reactionID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ReactionType - reactionType"))
}

// ReactionKEGGIDs is the resolver for the reactionKEGGIDs field.
func (r *queryResolver) ReactionKEGGIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ReactionKEGGIDs - reactionKEGGIDs"))
}

// ReactionKeggid is the resolver for the reactionKEGGID field.
func (r *queryResolver) ReactionKeggid(ctx context.Context, reactionID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ReactionKeggid - reactionKEGGID"))
}

// CompoundKeggid is the resolver for the compoundKEGGID field.
func (r *queryResolver) CompoundKeggid(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundKeggid - compoundKEGGID"))
}

// CompoundKEGGIDs is the resolver for the compoundKEGGIDs field.
func (r *queryResolver) CompoundKEGGIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundKEGGIDs - compoundKEGGIDs"))
}

// ChemicalFormulas is the resolver for the chemicalFormulas field.
func (r *queryResolver) ChemicalFormulas(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: ChemicalFormulas - chemicalFormulas"))
}

// ChemicalFormula is the resolver for the chemicalFormula field.
func (r *queryResolver) ChemicalFormula(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ChemicalFormula - chemicalFormula"))
}

// CompoundCASNumbers is the resolver for the compoundCASNumbers field.
func (r *queryResolver) CompoundCASNumbers(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundCASNumbers - compoundCASNumbers"))
}

// CompoundCASNumber is the resolver for the compoundCASNumber field.
func (r *queryResolver) CompoundCASNumber(ctx context.Context, compoundID *string) (*string, error) {
	panic(fmt.Errorf("not implemented: CompoundCASNumber - compoundCASNumber"))
}

// CompoundIDByFormula is the resolver for the compoundIDByFormula field.
func (r *queryResolver) CompoundIDByFormula(ctx context.Context, formula *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundIDByFormula - compoundIDByFormula"))
}

// CompoundNameBySearchTerm is the resolver for the compoundNameBySearchTerm field.
func (r *queryResolver) CompoundNameBySearchTerm(ctx context.Context, searchTerm *string) ([]*string, error) {
	panic(fmt.Errorf("not implemented: CompoundNameBySearchTerm - compoundNameBySearchTerm"))
}

// ModedIDByFileName is the resolver for the modedIDByFileName field.
func (r *queryResolver) ModedIDByFileName(ctx context.Context, fileName *string) (*string, error) {
	panic(fmt.Errorf("not implemented: ModedIDByFileName - modedIDByFileName"))
}

// FbaModelIDs is the resolver for the fbaModelIDs field.
func (r *queryResolver) FbaModelIDs(ctx context.Context) ([]*string, error) {
	panic(fmt.Errorf("not implemented: FbaModelIDs - fbaModelIDs"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
