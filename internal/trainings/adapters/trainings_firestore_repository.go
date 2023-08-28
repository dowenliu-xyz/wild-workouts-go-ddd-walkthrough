package adapters

import (
	"context"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/auth"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app"
)

type TrainingModel struct {
	UUID     string `firestore:"Uuid"`
	UserUUID string `firestore:"UserUuid"`
	User     string `firestore:"User"`

	Time  time.Time `firestore:"Time"`
	Notes string    `firestore:"Notes"`

	ProposedTime   *time.Time `firestore:"ProposedTime"`
	MoveProposedBy *string    `firestore:"MoveProposedBy"`
}

type TrainingsFirestoreRepository struct {
	firestoreClient *firestore.Client
}

func NewTrainingsFirestoreRepository(
	firestoreClient *firestore.Client,
) TrainingsFirestoreRepository {
	return TrainingsFirestoreRepository{
		firestoreClient: firestoreClient,
	}
}

func (d TrainingsFirestoreRepository) trainingsCollection() *firestore.CollectionRef {
	return d.firestoreClient.Collection("trainings")
}

func (d TrainingsFirestoreRepository) AllTrainings(ctx context.Context) ([]app.Training, error) {
	query := d.trainingsCollection().Query.Where("Time", ">=", time.Now().Add(-time.Hour*24))

	iter := query.Documents(ctx)

	return trainingModelsToApp(iter)
}

func (d TrainingsFirestoreRepository) FindTrainingsForUser(ctx context.Context, user auth.User) ([]app.Training, error) {
	query := d.trainingsCollection().Query.
		Where("Time", ">=", time.Now().Add(-time.Hour*24)).
		Where("UserUuid", "==", user.UUID)

	iter := query.Documents(ctx)

	return trainingModelsToApp(iter)
}

func trainingModelsToApp(iter *firestore.DocumentIterator) ([]app.Training, error) {
	var trainings []app.Training

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		t := TrainingModel{}
		if err := doc.DataTo(&t); err != nil {
			return nil, errors.WithStack(err)
		}

		trainings = append(trainings, app.Training(t)) // 内存布局相同，可以直接强转
	}

	sort.Slice(trainings, func(i, j int) bool { return trainings[i].Time.Before(trainings[j].Time) })

	return trainings, nil
}

func (d TrainingsFirestoreRepository) CreateTraining(ctx context.Context, training app.Training, createFn func() error) error {
	collection := d.trainingsCollection()

	trainingModel := TrainingModel(training) // 内存布局相同，直接强转

	// TODO pkg/errors Wrap/WithStack ? 问题是闭包内可能也是 Wrap/WithStack 的错误，这样会导致堆栈信息重复
	// 目前看，没有办法判断 error 是否是被 pkg/errors 包 Wrap/WithStack 处理过的，感觉需要 pkg/errors 提供一个 HasStack 方法，判断是否已经有堆栈信息了，这样如果有就不再加了

	return errors.WithStack(d.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		docs, err := tx.Documents(collection.Where("Time", "==", trainingModel.Time)).GetAll()
		if err != nil {
			return errors.Wrap(err, "unable to get actual docs")
		}
		if len(docs) > 0 {
			return errors.Errorf("there is training already at %s", trainingModel.Time)
		}

		err = createFn()
		if err != nil {
			return err
		}

		return errors.WithStack(tx.Create(collection.Doc(trainingModel.UUID), trainingModel))
	}))
}

func (d TrainingsFirestoreRepository) CancelTraining(ctx context.Context, trainingUUID string, deleteFn func(app.Training) error) error {
	trainingsCollection := d.trainingsCollection()

	return errors.WithStack(d.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		trainingDocumentRef := trainingsCollection.Doc(trainingUUID)

		firestoreTraining, err := tx.Get(trainingDocumentRef)
		if err != nil {
			return errors.Wrap(err, "unable to get actual docs")
		}

		training := TrainingModel{}
		err = firestoreTraining.DataTo(&training)
		if err != nil {
			return errors.Wrap(err, "unable to load document")
		}

		err = deleteFn(app.Training(training)) // 内存布局相同，直接强转
		if err != nil {
			return err
		}

		return errors.WithStack(tx.Delete(trainingDocumentRef))
	}))
}

func (d TrainingsFirestoreRepository) RescheduleTraining(
	ctx context.Context,
	trainingUUID string,
	newTime time.Time,
	updateFn func(app.Training) (app.Training, error),
) error {
	collection := d.trainingsCollection()

	return errors.WithStack(d.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(d.trainingsCollection().Doc(trainingUUID))
		if err != nil {
			return errors.Wrap(err, "could not find training")
		}

		docs, err := tx.Documents(collection.Where("Time", "==", newTime)).GetAll()
		if err != nil {
			return errors.Wrap(err, "unable to get actual docs")
		}
		if len(docs) > 0 {
			return errors.Errorf("there is training already at %s", newTime)
		}

		var training TrainingModel
		err = doc.DataTo(&training)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal training")
		}

		updatedTraining, err := updateFn(app.Training(training)) // 内存布局相同，直接强转
		if err != nil {
			return err
		}

		return errors.WithStack(tx.Set(collection.Doc(training.UUID), TrainingModel(updatedTraining))) // 内存布局相同，直接强转
	}))
}

func (d TrainingsFirestoreRepository) ApproveTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(app.Training) (app.Training, error)) error {
	return errors.WithStack(d.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(d.trainingsCollection().Doc(trainingUUID))
		if err != nil {
			return errors.Wrap(err, "could not find training")
		}

		var training TrainingModel
		err = doc.DataTo(&training)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal training")
		}

		updatedTraining, err := updateFn(app.Training(training)) // 内存布局相同，直接强转
		if err != nil {
			return err
		}

		return errors.WithStack(tx.Set(d.trainingsCollection().Doc(training.UUID), TrainingModel(updatedTraining))) // 内存布局相同，直接强转
	}))
}

func (d TrainingsFirestoreRepository) RejectTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(app.Training) (app.Training, error)) error {
	return errors.WithStack(d.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(d.trainingsCollection().Doc(trainingUUID))
		if err != nil {
			return errors.Wrap(err, "could not find training")
		}

		var training TrainingModel
		err = doc.DataTo(&training)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal training")
		}

		updatedTraining, err := updateFn(app.Training(training)) // 内存布局相同，直接强转
		if err != nil {
			return err
		}

		return errors.WithStack(tx.Set(d.trainingsCollection().Doc(training.UUID), TrainingModel(updatedTraining))) // 内存布局相同，直接强转
	}))
}
